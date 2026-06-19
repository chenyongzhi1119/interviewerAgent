package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"interviewer-agent/internal/agent"
	"interviewer-agent/internal/llm"
	"interviewer-agent/internal/server"
)

//go:embed all:web
var webFiles embed.FS

//go:embed companies
var companiesEmbed embed.FS

func main() {
	// 托盘模式：-tray 参数，或 .app 包内自动检测（无终端环境）
	trayMode := len(os.Args) > 1 && os.Args[1] == "-tray"
	if !trayMode {
		// 如果父进程不是终端（例如从 .app 双击启动），自动切换托盘模式
		if os.Getenv("TERM") == "" && os.Getenv("TERM_PROGRAM") == "" {
			trayMode = true
		}
	}

	if trayMode {
		// 单实例控制：如果服务已在运行，直接打开浏览器，本进程退出
		if isInterviewProRunning() {
			openBrowserURL("http://localhost:8080")
			return
		}
		runAsTrayApp(func(ready chan<- struct{}) {
			startHTTPServer(ready)
		})
		return
	}

	// 普通 CLI 模式
	startHTTPServer(nil)
}

// startHTTPServer 启动 HTTP 服务器。
// 若 ready 不为 nil，服务就绪后向其发送信号（用于托盘模式自动打开浏览器）。
func startHTTPServer(ready chan<- struct{}) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 检测是否在 .app 包内运行
	runningAsApp := isRunningAsApp()

	companiesDir := os.Getenv("COMPANIES_DIR")
	if companiesDir == "" {
		if runningAsApp {
			// 从 embed 提取 companies 到临时目录
			var err error
			companiesDir, err = extractEmbeddedCompanies()
			if err != nil {
				log.Fatalf("解压公司模板失败: %v", err)
			}
		} else {
			companiesDir = "companies"
		}
	}

	sessionsDir := os.Getenv("SESSIONS_DIR")
	if sessionsDir == "" {
		if runningAsApp {
			// 使用标准 macOS 数据目录
			home, _ := os.UserHomeDir()
			sessionsDir = filepath.Join(home, "Library", "Application Support", "InterviewPro", "sessions")
		} else {
			sessionsDir = "sessions"
		}
	}

	// ── 注册 LLM 供应商 ────────────────────────────────────────
	providers := llm.Registry{}
	if key := os.Getenv("ANTHROPIC_API_KEY"); key != "" {
		providers["anthropic"] = llm.NewAnthropic(key, os.Getenv("ANTHROPIC_MODEL"))
		log.Println("Provider enabled: Claude (Anthropic)")
	}
	if key := os.Getenv("OPENAI_API_KEY"); key != "" {
		providers["openai"] = llm.NewOpenAICompat("openai", key, os.Getenv("OPENAI_BASE_URL"), os.Getenv("OPENAI_MODEL"))
		log.Println("Provider enabled: GPT (OpenAI)")
	}
	if key := os.Getenv("DEEPSEEK_API_KEY"); key != "" {
		providers["deepseek"] = llm.NewOpenAICompat("deepseek", key, os.Getenv("DEEPSEEK_BASE_URL"), os.Getenv("DEEPSEEK_MODEL"))
		log.Println("Provider enabled: DeepSeek")
	}
	if key := os.Getenv("GLM_API_KEY"); key != "" {
		providers["glm"] = llm.NewOpenAICompat("glm", key, os.Getenv("GLM_BASE_URL"), os.Getenv("GLM_MODEL"))
		log.Println("Provider enabled: GLM (智谱 AI)")
	}
	if key := os.Getenv("QWEN_API_KEY"); key != "" {
		providers["qwen"] = llm.NewOpenAICompat("qwen", key, os.Getenv("QWEN_BASE_URL"), os.Getenv("QWEN_MODEL"))
		log.Println("Provider enabled: Qwen (阿里云)")
	}

	// 托盘模式允许零供应商启动（用户在 UI 里填 key）
	if len(providers) == 0 && ready == nil {
		log.Fatal("No LLM provider configured. Set at least one of: ANTHROPIC_API_KEY, OPENAI_API_KEY, DEEPSEEK_API_KEY, GLM_API_KEY, QWEN_API_KEY")
	}

	store, err := agent.NewStore(providers, companiesDir, sessionsDir)
	if err != nil {
		log.Fatalf("failed to init store: %v", err)
	}

	var webFS fs.FS
	if customWeb := os.Getenv("WEB_DIR"); customWeb != "" {
		webFS = os.DirFS(customWeb)
	} else {
		var subErr error
		webFS, subErr = fs.Sub(webFiles, "web")
		if subErr != nil {
			log.Fatalf("failed to embed web assets: %v", subErr)
		}
	}

	h := server.New(store, webFS)
	addr := fmt.Sprintf(":%s", port)
	log.Printf("面试官 Agent 已启动 → http://localhost%s", addr)

	// 服务就绪信号（供托盘模式打开浏览器）
	if ready != nil {
		ready <- struct{}{}
	}

	if err := http.ListenAndServe(addr, h); err != nil {
		log.Fatal(err)
	}
}

// isRunningAsApp 检测是否在 macOS .app 包内运行（双击启动）。
func isRunningAsApp() bool {
	exe, err := os.Executable()
	if err != nil {
		return false
	}
	return strings.Contains(exe, ".app/Contents/MacOS/")
}

// extractEmbeddedCompanies 将 embed 的 companies/*.yaml 解压到临时目录。
func extractEmbeddedCompanies() (string, error) {
	tmpDir, err := os.MkdirTemp("", "interviewpro-companies-*")
	if err != nil {
		return "", err
	}
	entries, err := companiesEmbed.ReadDir("companies")
	if err != nil {
		return "", err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		data, err := companiesEmbed.ReadFile("companies/" + e.Name())
		if err != nil {
			return "", err
		}
		if err := os.WriteFile(filepath.Join(tmpDir, e.Name()), data, 0644); err != nil {
			return "", err
		}
	}
	return tmpDir, nil
}

// isInterviewProRunning 检测 8080 端口是否已有 InterviewPro 实例在运行。
func isInterviewProRunning() bool {
	client := &http.Client{Timeout: 800 * time.Millisecond}
	resp, err := client.Get("http://localhost:8080/api/companies")
	if err != nil {
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == 200
}

// openBrowserURL 在 macOS 上用默认浏览器打开 URL。
func openBrowserURL(url string) {
	exec.Command("open", url).Start()
}

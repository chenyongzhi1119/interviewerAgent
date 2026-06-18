package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"interviewer-agent/internal/agent"
	"interviewer-agent/internal/llm"
	"interviewer-agent/internal/server"
)

//go:embed all:web
var webFiles embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	companiesDir := os.Getenv("COMPANIES_DIR")
	if companiesDir == "" {
		companiesDir = "companies"
	}
	sessionsDir := os.Getenv("SESSIONS_DIR")
	if sessionsDir == "" {
		sessionsDir = "sessions"
	}

	// ── Register available LLM providers ────────────────────────────────────
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

	if len(providers) == 0 {
		log.Fatal("No LLM provider configured. Set at least one of: ANTHROPIC_API_KEY, OPENAI_API_KEY, DEEPSEEK_API_KEY, GLM_API_KEY, QWEN_API_KEY")
	}

	store, err := agent.NewStore(providers, companiesDir, sessionsDir)
	if err != nil {
		log.Fatalf("failed to init store: %v", err)
	}

	var webFS fs.FS
	if customWeb := os.Getenv("WEB_DIR"); customWeb != "" {
		webFS = os.DirFS(customWeb)
		log.Printf("Serving frontend from filesystem: %s", customWeb)
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
	if err := http.ListenAndServe(addr, h); err != nil {
		log.Fatal(err)
	}
}

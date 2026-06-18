//go:build darwin

package main

import (
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/getlantern/systray"
)

func isPortInUse(addr string) bool {
	conn, err := net.DialTimeout("tcp", "localhost"+addr, 500*time.Millisecond)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// runAsTrayApp 启动菜单栏模式：
//  1. 在后台启动 HTTP 服务器（传入 serverReady channel）
//  2. 服务就绪后自动打开浏览器
//  3. 菜单栏图标提供「打开」和「退出」
func runAsTrayApp(startServer func(ready chan<- struct{})) {
	serverReady := make(chan struct{}, 1)

	systray.Run(
		func() { onTrayReady(serverReady, startServer) },
		func() { os.Exit(0) },
	)
}

func onTrayReady(serverReady chan struct{}, startServer func(ready chan<- struct{})) {
	systray.SetIcon(trayIconPNG())
	systray.SetTooltip("大厂面试官 Agent")

	mOpen := systray.AddMenuItem("打开 InterviewPro", "在浏览器中打开")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出 InterviewPro")

	// 尝试启动服务器（若端口已被占用则直接复用已有实例）
	go func() {
		if isPortInUse(":8080") {
			// 端口已被占用（可能是另一个实例），直接复用
			serverReady <- struct{}{}
		} else {
			startServer(serverReady)
		}
	}()

	// 服务就绪后自动打开浏览器
	go func() {
		<-serverReady
		time.Sleep(400 * time.Millisecond)
		openBrowser()
	}()

	// 处理菜单点击
	go func() {
		for {
			select {
			case <-mOpen.ClickedCh:
				openBrowser()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func openBrowser() {
	exec.Command("open", "http://localhost:8080").Start()
}

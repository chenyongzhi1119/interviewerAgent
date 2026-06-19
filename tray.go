//go:build darwin

package main

import (
	"net"
	"os"
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

// runAsTrayApp 启动菜单栏托盘模式。
func runAsTrayApp(startServer func(ready chan<- struct{})) {
	serverReady := make(chan struct{}, 1)
	systray.Run(
		func() { onTrayReady(serverReady, startServer) },
		func() { os.Exit(0) },
	)
}

func onTrayReady(serverReady chan struct{}, startServer func(ready chan<- struct{})) {
	systray.SetIcon(trayIconPNG())
	systray.SetTooltip("InterviewPro · 大厂面试官 Agent")

	mOpen := systray.AddMenuItem("打开 InterviewPro", "在浏览器中打开")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出 InterviewPro", "停止服务并完全退出")

	// 后台启动服务器
	go startServer(serverReady)

	// 服务就绪后自动打开浏览器
	go func() {
		<-serverReady
		time.Sleep(400 * time.Millisecond)
		openBrowserURL("http://localhost:8080")
	}()

	// 处理菜单点击
	go func() {
		for {
			select {
			case <-mOpen.ClickedCh:
				openBrowserURL("http://localhost:8080")
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

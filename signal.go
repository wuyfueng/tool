package tools

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitExit 等待退出
func WaitExit() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
}

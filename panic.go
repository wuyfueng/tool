package tools

import (
	"log"
	"runtime/debug"
)

// Async 异步处理并捕获panic
func Async(f func()) {
	go func() {
		Do(f, func(pan *Panic) {
			log.Printf("goroutine panic: %v, Stack: %v", pan.Reason, debug.Stack())
		})
	}()
}

package tool

import (
	"fmt"
	"time"
)

// AsyncDelayExec 异步延迟执行
func AsyncDelayExec(f func(), execTime time.Time) {
	go func() {
		do := func() {
			for now := range time.Tick(time.Second) {
				if now.After(execTime) {
					f()
					return
				}
			}
		}

		Do(do, func(pan *Panic) {
			fmt.Println(fmt.Sprintf("AsyncDelayExec goroutine panic: %v", pan.Reason))
		})
	}()
}

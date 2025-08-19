package homework2

import (
	"fmt"
	"sync"
	"time"
)

func Schedular(tasks []func()) {
	/**
	wg.Add(1) 的意思是将 sync.WaitGroup 的计数器加 1，表示有一个新的 goroutine 需要等待。每启动一个 goroutine 前都要调用一次 Add(1)，
	等 goroutine 完成后再调用 wg.Done()，这样主线程可以通过 wg.Wait() 等待所有 goroutine 结束。
	*/
	var wg sync.WaitGroup
	// 使用 WaitGroup 来等待所有 goroutine 完成
	for index, task := range tasks {
		wg.Add(1)
		go func(index int, task func()) {
			defer wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			fmt.Println("Task", index, "completed in", duration.String())
		}(index, task)

	}
	wg.Wait()
}

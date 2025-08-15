package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*goroutine 是轻量线程 协程，创建一个 goroutine 所需的资源开销很小，所以可以创建非常多的 goroutine 来并发工作。

它们是由 Go 运行时调度的。调度过程就是 Go 运行时把 goroutine 任务分配给 CPU 执行的过程。

但是 goroutine 不是通常理解的线程，线程是操作系统调度的。

GO中如果想让某个任务并发或者异步执行，只需要把任务封装成为一个函数或闭包，交给 goroutine执行即可
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

func main_go() {

	go func() {
		fmt.Println("run goroutine in closure")
	}()
	// 最后的小括号里的参数 "gorouine :closure with param" 是传递给匿名函数的参数
	go func(s string) {
		fmt.Println("run goroutine in closure with param:", s)
	}("gorouine :closure with param")

	go say("in goroutine:world")
	say("hello")

}

//go 也有线程安全问题

type SafeCounter struct {
	mu    sync.Mutex // 互斥锁
	count int
}

func (s *SafeCounter) Increment() {
	s.mu.Lock()         // 上锁
	defer s.mu.Unlock() // 结束解锁
	s.count++
}

func (s *SafeCounter) GetCount() int {
	s.mu.Lock()         // 上锁
	defer s.mu.Unlock() // 结束解锁
	return s.count
}

type UnSafeCounter struct {
	count int
}

func (u *UnSafeCounter) Increment() {
	u.count++
}

func (u *UnSafeCounter) GetCount() int {
	return u.count
}

func main_lock() {
	//counter := &SafeCounter{}
	counter := &UnSafeCounter{}

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				counter.Increment()
			}
		}()
	}

	//等待一段时间
	time.Sleep(2 * time.Second)

	fmt.Println("SafeCounter count:", counter.GetCount())

}

/**
// 发送数据
channel_name <- variable_name_or_value

// 接收数据
value_name, ok_flag := <- channel_name
value_name := <- channel_name

// 关闭channel
close(channel_name)


//仅发送数据
func <method_name>(<channel_name> chan <- <type>)

//仅接收数据
func <method_name>(<channel_name> <-chan <type>)
*/

// 接收数据 chan在后 ，从chan里发走  ch <- chan int
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Println("receiveOnly:", v)
	}
}

// 发送数据 chan在前 想chan里发
func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("sendOnly:", i)
	}
	close(ch) // 发送完数据后关闭通道

}

/*
*
大部分时候，流程是根据数据驱动的，channel 会被使用得更频繁。

channel 擅长的是**数据流动的场景**：

1. 传递数据的所有权，即把某个数据发送给其他协程。
2. 分发任务，每个任务都是一个数据。
3. 交流异步结果，结果是一个数据。

而锁使用的场景更偏向**同一时间只给一个协程访问数据的权限**：

1. 访问缓存
2. 管理状态
*/
func main18() {
	ch := make(chan int, 3)

	//启动发送
	go sendOnly(ch)

	//启动接收
	go receiveOnly(ch)

	//使用select进行多路复用

	timeout := time.After(2 * time.Second) // 等待一段时间，确保所有数据都被发送和接收完毕

	for {

		select {

		case v, ok := <-ch:
			if !ok {
				fmt.Println("channel is closed")
				return
			}
			fmt.Println("received value:", v)

		case <-timeout:
			fmt.Println("timeout")
			return

		default:
			fmt.Println("no data received, waiting...")
			time.Sleep(500 * time.Millisecond) // 等待一段时间再检查
		}
	}
}

package homework2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main22() {
	//1 指针
	//num := 10
	//homework2.Add(&num)
	//fmt.Println("num after add:", num) // 输出: num after add: 20
	//
	//arr := make([]int, 4, 10)
	//arr[0] = 1
	//arr[1] = 2
	//arr[2] = 3
	//arr[3] = 4
	//homework2.ChengArray(&arr)
	//fmt.Println("arr after chengArray:", arr) // 输出: arrs after

	//2 goroutine
	//var wg = sync.WaitGroup{}
	//wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//	for i := 1; i < 10; i += 2 {
	//		fmt.Println("奇数:", i)
	//	}
	//
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	for i := 2; i < 10; i += 2 {
	//		fmt.Println("偶数:", i)
	//	}
	//
	//}()
	//wg.Wait()
	//
	//tasks := []func(){
	//	func() { time.Sleep(500 * time.Millisecond) },
	//	func() { time.Sleep(200 * time.Millisecond) },
	//	func() { time.Sleep(300 * time.Millisecond) },
	//}
	//
	//homework2.Schedular(tasks)

	// 3 面向对象
	//r := homework2.Rectangle{Width: 3, Height: 4}
	//c := homework2.Circle{Radius: 5}
	//
	//fmt.Println("矩形面积:", r.Area())
	//fmt.Println("矩形周长:", r.Perimeter())
	//fmt.Println("圆面积:", c.Area())
	//fmt.Println("圆周长:", c.Perimeter())
	//
	//e := homework2.Employee{Person: homework2.Person{Name: "张三", Age: 18}, EmployeeID: "123456"}
	//fmt.Println(e.PrintInfo())

	//4 channel

	//ch := make(chan int)
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		ch <- i
	//	}
	//	close(ch)
	//}()
	//
	////消费
	//for num := range ch {
	//	fmt.Println("消费:", num)
	//}
	//
	//chBuffer := make(chan int, 100)
	//
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		chBuffer <- i
	//	}
	//	close(chBuffer)
	//}()
	//
	//for num := range chBuffer {
	//	fmt.Println("消费:", num)
	//}

	//锁

	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}

		}()
	}
	wg.Wait()
	fmt.Println("lock counter:", counter)

	//原子锁
	var counter2 int64
	var wg2 sync.WaitGroup

	wg2.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			//每个goroutine都需要调用Done来通知WaitGroup
			//否则WaitGroup不会等待所有goroutine完成
			defer wg2.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter2, 1)
			}

		}()

	}
	wg2.Wait()
	fmt.Println("atomic counter:", counter2)

}

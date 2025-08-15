package basecode

import (
	"fmt"
)

func main9() {

	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//b := 1
	//for b < 10 {
	//	fmt.Println(b)
	//	b++
	//}

	//3 无限循环
	//ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	//var stared bool
	//var stopped atomic.Bool
	//
	//for {
	//	if !stared {
	//		stared = true
	//		fmt.Println("start")
	//		go func() {
	//			for {
	//				select {
	//				case <-ctx.Done():
	//					fmt.Println("context done")
	//					stopped.Store(true)
	//					return
	//				}
	//			}
	//		}()
	//	}
	//	fmt.Println("main")
	//	if stopped.Load() {
	//		break
	//	}
	//}

	//遍历数组
	var a [10]string
	a[0] = "hello"
	for i, value := range a {
		fmt.Println(i, value)
	}

	//遍历切片

	s := make([]string, 10)

	s[0] = "hello"

	for i, value := range s {
		fmt.Println(i, value)
	}

	// 遍历map

	m := make(map[string]string)

	m["b"] = "hello,b"
	m["a"] = "hello,a"
	m["c"] = "hello,c"
	m["d"] = "hello,d"
	m["e"] = "hello,e"
	m["f"] = "hello,f"

	for key := range m {
		fmt.Println("key", key)
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

}

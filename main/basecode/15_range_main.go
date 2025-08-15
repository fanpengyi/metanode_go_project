package basecode

import (
	"fmt"
	"reflect"
	"time"
)

func main_range_str() {
	str1 := "abc123"
	for index, value := range str1 {
		fmt.Println(index, string(value), str1[index])
	}
	str2 := "测试中文"

	for index, value := range str2 {
		fmt.Println(index, string(value), str2[index])
	}

	//中文使用rune统计的是字符数，使用byte统计的是字节数
	runes := []rune(str2)
	bytes := []byte(str2)

	fmt.Printf("len(runesFromStr2) = %d\n", len(runes))
	fmt.Printf("len(bytes) = %d\n", len(bytes))

	str3 := "a1中文"
	for index, value := range str3 {
		//当直接使用下标取字符串某个下标位置上的值时，取出来的是 byte 值。
		fmt.Printf("str3 -- index:%d, index value:%d\n", index, str3[index])
		//而使用 for range 循环遍历字符串时，取出来的是 rune 值。 string(value) 可以将 rune 值转换为字符串。
		fmt.Printf("str3 -- index:%d, range value:%d,string value:%s \n", index, value, string(value))
	}

}

func main_range_array() {
	//array := [...]int{1, 2, 3}
	//slice := []int{4, 5, 6}
	//
	//// 方法1：只拿到数组的下标索引
	//for index := range array {
	//	fmt.Printf("array -- index=%d value=%d \n", index, array[index])
	//}
	//for index := range slice {
	//	fmt.Printf("slice -- index=%d value=%d \n", index, slice[index])
	//}
	//fmt.Println()
	//
	//// 方法2：同时拿到数组的下标索引和对应的值
	//for index, value := range array {
	//	fmt.Printf("array -- index=%d index value=%d \n", index, array[index])
	//	fmt.Printf("array -- index=%d range value=%d \n", index, value)
	//}
	//for index, value := range slice {
	//	fmt.Printf("slice -- index=%d index value=%d \n", index, slice[index])
	//	fmt.Printf("slice -- index=%d range value=%d \n", index, value)
	//}
	//fmt.Println()

	array := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	slice := [][]int{{1, 2}, {3}}
	// 只拿到行的索引
	for index := range array {
		// array[index]类型是一维数组
		fmt.Println(reflect.TypeOf(array[index]))
		fmt.Printf("array -- index=%d, value=%v\n", index, array[index])
	}

	for index := range slice {
		// slice[index]类型是一维数组
		fmt.Println(reflect.TypeOf(slice[index]))
		fmt.Printf("slice -- index=%d, value=%v\n", index, slice[index])
	}

	// 拿到行索引和该行的数据
	fmt.Println("print array element")
	for row_index, row_value := range array {
		fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
	}

	fmt.Println("print array slice")
	for row_index, row_value := range slice {
		fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
	}

	// 双重遍历，拿到每个元素的值
	for row_index, row_value := range array {
		for col_index, col_value := range row_value {
			fmt.Printf("array[%d][%d]=%d ", row_index, col_index, col_value)
		}
		fmt.Println()
	}
	for row_index, row_value := range slice {
		for col_index, col_value := range row_value {
			fmt.Printf("slice[%d][%d]=%d ", row_index, col_index, col_value)
		}
		fmt.Println()
	}

}

func addData(ch chan int) {
	size := cap(ch) // 获取通道的容量
	for i := 0; i < size; i++ {
		ch <- i
		time.Sleep(1 * time.Second) // 模拟数据添加的延时
	}
	close(ch)
}

func main_range_channel() {
	ch := make(chan int, 10)

	go addData(ch)
	// 使用 range 遍历通道
	for i := range ch {
		fmt.Println("Received:", i)
	}

}

func main15() {
	hash := map[string]int{
		"a": 1,
		"f": 2,
		"z": 3,
		"c": 4,
	}

	for key := range hash {
		fmt.Printf("key=%s, value=%d\n", key, hash[key])
	}

	for key, value := range hash {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
}

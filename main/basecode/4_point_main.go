package basecode

func main1() {
	// 1. 指针的基本使用
	//var p1 *int
	//var p2 *string
	//i := 1
	//s := "hello"
	//
	//p1 = &i
	//p2 = &s
	//
	//p3 := &p2
	//fmt.Println(p1)
	//fmt.Println(p2)
	//fmt.Println(p3)

	// 2指针和 unsafe.Pointer 转换
	//*T <---> unsafe.Pointer <--> uintptr
	//a := "hello world"
	//upA := uintptr(unsafe.Pointer(&a))
	//upA += 5 // 偏移5个字节
	//
	//c := (*uint8)(unsafe.Pointer(upA))
	//fmt.Println(*c)

}

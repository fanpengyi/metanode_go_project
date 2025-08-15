package basecode

import "fmt"

func main_slice_init() {
	//声明切片
	// 方式1，声明并初始化一个空的切片
	var s1 []int = []int{}
	fmt.Println(s1)

	// 方式2，类型推导，并初始化一个空的切片
	var s2 = []int{}
	fmt.Println(s2)
	// 方式3，与方式2等价
	s3 := []int{}
	fmt.Println(s3)
	// 方式4，与方式1、2、3 等价，可以在大括号中定义切片初始元素
	s4 := []int{1, 2, 3, 4}
	fmt.Println(s4)
	// 方式5，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为0
	s5 := make([]int, 0)
	fmt.Println(s5)
	// 方式6，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为2，指定容量参数4
	s6 := make([]int, 2, 4)
	fmt.Println(s6)
	// 方式7，引用一个数组，初始化切片
	arr := [5]int{6, 5, 4, 3, 2}
	fmt.Println(arr)
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := arr[2:]
	fmt.Println(s7)
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s8 := arr[1:3]
	fmt.Println(s8)
	// 从0到下标2的元素，创建一个新的切片
	s9 := arr[:2]
	fmt.Println(s9)

}

// 当切片是基于同一个数组创建的，那么对切片的修改会影响到原数组
func main13_slice_remove() {
	//a := [5]int{6, 5, 4, 3, 2}
	//// 从数组下标2开始，直到数组的最后一个元素
	//s7 := a[2:]
	//// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	//s8 := a[1:3]
	//// 从0到下标2的元素，创建一个新的切片
	//s9 := a[:2]
	//fmt.Println(s7)
	//fmt.Println(s8)
	//fmt.Println(s9)
	//a[0] = 9
	//a[1] = 8
	//a[2] = 7
	//fmt.Println(s7)
	//fmt.Println(s8)
	//fmt.Println(s9)

	//向指定位置添加元素
	ints := []int{1, 2, 3, 4, 5}
	//像下标为2的位置添加元素6
	/**
	ints[:2] 取出原切片前两个元素。
	append([]int{6}, ints[2:]...) 先构造一个只包含6的新切片，再把原切片下标2及之后的元素拼接到后面。
	最外层 append 把前两个元素和插入后的部分拼接起来，形成新的切片。
	最终效果是：在原切片下标2的位置插入了6，后面的元素依次后移。
	*/
	ints = append(ints[:2], append([]int{6}, ints[2:]...)...)
	fmt.Println("ints:", ints)

	//移除指定位置的元素
	s5 := []int{1, 2, 3, 5, 4}
	//移除下标为3的元素
	//s5[:3] 取出原切片前3个元素。
	//s5[4:] 取出原切片下标4及之后的元素。
	//append(s5[:3], s5[4:]...) 把前3个元素和下标4及之后的元素拼接起来，形成新的切片。
	//最终效果是：移除了原切片下标3的元素，后面的元素依次前移。
	s5 = append(s5[:3], s5[4:]...)
	fmt.Println("s5 = ", s5)
}

func main_slice_copy() {
	//复制切片 用内置函数 `copy()` 把某个切片中的所有元素复制到另一个切片，复制的长度是它们中最短的切片长度。
	//	src1 := []int{1, 2, 3}
	//	dst1 := make([]int, 4, 5)
	//
	//	src2 := []int{1, 2, 3, 4, 5}
	//	dst2 := make([]int, 3, 3)
	//
	//	fmt.Println("before copy, src1 = ", src1)
	//	fmt.Println("before copy, dst1 = ", dst1)
	//
	//	fmt.Println("before copy, src2 = ", src2)
	//	fmt.Println("before copy, dst2 = ", dst2)
	//
	//	copy(dst1, src1)
	//	copy(dst2, src2)
	//
	//	fmt.Println("after copy, src1 = ", src1)
	//	fmt.Println("after copy, dst1 = ", dst1)
	//
	//	fmt.Println("after copy, src2 = ", src2)
	//	fmt.Println("after copy, dst2 = ", dst2)
	//
	//切片类型实际上是比较特殊的指针类型，当声明一个切片类型时，就是声明了一个指针。
	//切片是一个指针类型
	//与数组不同的是，方法参数传切片值类型，修改是会影响外部实参的！
	//在不使用append函数情况下，在函数内部对切片进行修改，外部实参也会被修改。
	s := make([]int, 3, 6)
	fmt.Println("s length:", len(s))
	fmt.Println("s capacity:", cap(s))
	fmt.Println("initial, s = ", s)
	s[1] = 2
	fmt.Println("set position 1, s = ", s)

	modifySlice(s)
	fmt.Println("after modifySlice, s = ", s)
}

func modifySlice(param []int) {
	param[0] = 1024
}

func main_slice_no_expand() {
	//不触发切片扩容场景
	//对原切片的修改会影响到新切片的值
	s := make([]int, 3, 6)
	fmt.Println("initial, s =", s)
	s[1] = 2
	fmt.Println("after set position 1, s =", s)

	s2 := append(s, 4)
	fmt.Println("after append, s2 length:", len(s2))
	fmt.Println("after append, s2 capacity:", cap(s2))
	fmt.Println("after append, s =", s)
	fmt.Println("after append, s2 =", s2)
	//原切片改值
	//s2 也会改
	s[0] = 1024
	fmt.Println("after set position 0, s =", s)
	fmt.Println("after set position 0, s2 =", s2)

	appendInFuncNoExpand(s)
	fmt.Println("after append in func, s =", s)
	fmt.Println("after append in func, s2 =", s2)
}

func appendInFuncNoExpand(param []int) {
	param = append(param, 1022)
	fmt.Println("in func, param =", param)
	param[2] = 512
	fmt.Println("set position 2 in func, param =", param)
}

func main13() {
	//对原切片添加元素 触发扩容
	s := make([]int, 2, 2)
	fmt.Println("initial, s =", s)

	s2 := append(s, 4)
	fmt.Println("after append, s length:", len(s))
	fmt.Println("after append, s capacity:", cap(s))

	fmt.Println("after append, s2 length:", len(s2))
	fmt.Println("after append, s2 capacity:", cap(s2))
	fmt.Println("after append, s =", s)
	fmt.Println("after append, s2 =", s2)

	s[0] = 1024
	fmt.Println("after set position 0, s =", s)
	fmt.Println("after set position 0, s2 =", s2)

	appendInFuncExpand(s2)
	fmt.Printf("s2指针地址: %p\n", &s2)
	fmt.Println("after append in func, s2 =", &s2[0])
	fmt.Println("after append in func, s2 =", s2)
}

func appendInFuncExpand(param []int) {
	// 直接修改会影响外部实参
	//param[0] = 11111

	//使用append函数不会影响外部实参
	//param和s2中元素存储的内存地址相同，
	//但是param和s2的指针不同
	//param指针地址: 0xc000092228   s2指针地址: 0xc0000921b0
	param = append(param, 511)
	fmt.Printf("param指针地址: %p\n", &param)
	//param2 := append(param1, 512)
	//fmt.Println("in func, param1 =", param1)
	fmt.Println("in func, param =", param)
	fmt.Println("in func, param =", &param[0])

	//fmt.Println("set position 2 in func, param2 =", param2)
}

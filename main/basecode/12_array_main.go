package basecode

import "fmt"

func main12_init() {
	//声明
	var a [5]int
	fmt.Println("a:", a)
	//赋值
	var maparr [2]map[string]string
	//map的零值是 nil
	fmt.Println("maparr:", maparr)

	//声明和初始化
	//var b [5]int = [5]int{1, 2, 3, 4, 5}

	//类型推导声明
	ints := [5]int{1, 2, 3, 4, 5}
	fmt.Println("ints:", ints)

	//使用 ...代替数组长度
	strings := [...]string{"a", "b", "c", "d"}
	fmt.Println("strings:", strings)

	//声明初始化指定下标
	position := [5]string{1: "position1", 3: "position3"}
	fmt.Println("position:", position)

}

// 访问
func main12_access() {

	//ints := [6]int{1, 2, 3, 4, 5, 6}
	//
	//for index, value := range ints {
	//	fmt.Println("index:", index, "value:", value)
	//}
	//
	//for i := 0; i < len(ints); i++ {
	//	fmt.Println("index:", i, "value:", ints[i])
	//}

	// 二维数组
	a := [3][2]int{
		{0, 1},
		{2, 3},
		{4, 5},
	}
	fmt.Println("a = ", a)

	// 三维数组
	b := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}
	fmt.Println("b = ", b)

	// 也可以省略各个位置的初始化,在后续代码中赋值
	c := [3][3][3]int{}
	c[2][2][1] = 5
	c[1][2][1] = 4
	fmt.Println("c = ", c)

}

type Custom struct {
	i int
}

var carr [5]*Custom = [5]*Custom{
	{6},
	{7},
	{8},
	{9},
	{10},
}

func main12_array() {
	a := [5]int{5, 4, 3, 2, 1}
	fmt.Println("before all, a = ", a)
	for i := range carr {
		fmt.Printf("before in main func, carr[%d] = %p, value = %v \n", i, &carr[i], *carr[i])
	}
	//值类型不会改变外部实参
	//printFuncParam(carr)

	//指针类型才可以改变外部实参
	printFuncParamPointer(&carr)

	for i := range carr {
		fmt.Printf("after in main func, carr[%d] = %p, value = %v \n", i, &carr[i], *carr[i])
	}

	//数组传递值类型 不会改变外部实参的值
	receiveArray(a)
	fmt.Println("after receiveArray, a = ", a)

	//数组改变外部实参的值 也需要传递指针类型
	receiveArrayPointer(&a)
	fmt.Println("after receiveArrayPointer, a = ", a)
}

func receiveArray(param [5]int) {
	fmt.Println("in receiveArray func, before modify, param = ", param)
	param[1] = -5
	fmt.Println("in receiveArray func, after modify, param = ", param)
}

func receiveArrayPointer(param *[5]int) {
	fmt.Println("in receiveArrayPointer func, before modify, param = ", param)
	param[1] = -5
	fmt.Println("in receiveArrayPointer func, after modify, param = ", param)
}

func printFuncParamPointer(param *[5]*Custom) {
	//for i := range param {
	//	fmt.Printf("in printFuncParamPointer func, param[%d] = %p, value = %v \n", i, &param[i], *param[i])
	//}
	fmt.Println("in printFuncParamPointer func, before modify, param = ", param)
	param[1] = &Custom{100}
	fmt.Println("in printFuncParamPointer func, after modify, param = ", param)

}

func printFuncParam(param [5]*Custom) {
	//for i := range param {
	//	fmt.Printf("in printFuncParamPointer func, param[%d] = %p, value = %v \n", i, &param[i], *param[i])
	//}
	fmt.Println("in printFuncParamPointer func, before modify, param = ", param)
	param[1] = &Custom{100}
	fmt.Println("in printFuncParamPointer func, after modify, param = ", param)

}

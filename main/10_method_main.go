package main

import "fmt"

type AType struct {
	i int
}

// 定义方法,方法包含接受者和方法名
func (a *AType) add(v int) int {
	a.i += v
	return a.i
}

// 声明函数 函数名 函数关键字 func
var function1 func(int) int

//func(a A) int

// 声明闭包
var square2 func(int) int = func(p int) int {
	p *= p
	return p
}

var method2 func(int) int = func(m int) int {
	m *= m
	return m
}

func main10() {
	a2 := AType{1}
	//把方法赋值给函数变量
	function1 = a2.add

	//声明一个闭包
	//返回值是一个函数
	returnFunc := func() func(int, string) (int, string) {
		fmt.Println("this is a anonymous function")
		return func(i int, s string) (int, string) {
			return i, s
		}
	}()
	//调用闭包函数
	i2, s := returnFunc(1, "hello")
	fmt.Println("call returnFunc:", i2, s)
	fmt.Println(a2.i)
	fmt.Println("after call function1: a2.i", function1(10))
	fmt.Println(a2.i)

	a2.add(20)
	fmt.Println("after call a2.add: a2.i", a2.i)

}

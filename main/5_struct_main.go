package main

import "fmt"

type A struct {
	a string
}

/*
不管接受者是什么类型都不影响变量调用方法。但是他们调用方法后，产生的结果会有一些不同。

值接受者（Value Receiver）: 值接受者的方法操作的是值的副本，在这个方法中可以随意修改值接受者的字段的值，但不会影响原始实例。

指针接受者（Pointer Receiver）：指针接受者的方法操作的是原实例的指针，修改指针接受者的任意字段，也意味着修改了原实例。


当结构体比较复杂并且数据量比较大时（比如持有一个[]byte，长度有 8000 以上），尽量使用指针接受者。因为执行值接受者方法时，会创建一个实例的副本，会占用比较大的栈空间，即使栈空间的清理很容易，但在一些资源受限的平台上运行，可能会导致内存爆掉，而被系统强制杀死进程。

当结构体只是存储数据，并且执行方法后，不希望方法中的临时修改影响到原来的实例，那么就可以选择使用值接受者。

*/

// 方法有接收器 一般是结构体中的叫方法 (a A)
func (a A) string() string {
	return a.a
}

func (a A) stringA() string {
	return a.a
}

func (a A) setA(v string) {
	a.a = v
}

func (a *A) stringPA() string {
	return a.a
}

func (a *A) setPA(v string) {
	a.a = v
}

type B struct {
	A
	b string
}

func (b B) string() string {
	return b.b
}

func (b B) stringB() string {
	return b.b
}

type C struct {
	B
	a string
	b string
	c string
	d []byte
}

func (c C) string() string {
	return c.c
}

func (c C) modityD() {
	c.d[2] = 3
}

func callStructMethod() {
	var a A
	a = A{
		a: "a",
	}
	a.string()
}

// 构造函数
func NewC() C {
	return C{
		B: B{
			A: A{
				a: "ba",
			},
			b: "b",
		},
		a: "ca",
		b: "cb",
		c: "c",
		d: []byte{1, 2, 8},
	}
}

// 1 GO 没有方法重写，不允许相同名称的方法存在，即使参数不同
// 2 结构体类型变量和结构体指针变量都可以直接访问结构体声明的字段和调用声明的方法
// 3 没有构造函数，构造函数自己定义，一般按照习惯 定义的名称是New<structName>的函数来实例化 一个比较复杂的结构体 返回结构体类型
func main5() {
	c := NewC()
	cp := &c
	fmt.Println(c.string())
	//调用的是A的方法 没有重写
	fmt.Println(c.stringA())
	fmt.Println(c.stringB())

	fmt.Println(cp.string())
	fmt.Println(cp.stringA())
	fmt.Println(cp.stringB())

	//调研结构体类型参数的赋值 不会改变外部实参的值
	c.setA("1a")
	fmt.Println("------------------c.setA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	cp.setA("2a")
	fmt.Println("------------------cp.setA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	//调用指针类型的赋值方法 会修改外部实参
	c.setPA("3a")
	fmt.Println("------------------c.setPA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	cp.setPA("4a")
	fmt.Println("------------------cp.setPA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)
	//调用结构体类型的 如果是数组类型会改变实参
	fmt.Println("before ------------------cp.modityD")
	fmt.Println(cp.d)
	cp.modityD()
	fmt.Println(" after------------------cp.modityD")
	fmt.Println(cp.d)

}

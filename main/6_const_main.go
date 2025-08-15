package main

import "fmt"

const a = 100

const (
	h    byte = 3
	i         = "value"
	j, k      = "v", 4
)

// 定义枚举类型 ，类型别名，让常量定义的枚举类型的作用显得更直观
type Gender string

const (
	Male   = "Male"
	Female = "Female"
)

// 没有 struct 也可以定义接收器 *Gender ，方法名 String
func (g *Gender) String() string {
	switch *g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknow"
	}
}

func (g *Gender) IsMale() bool {
	return *g == Male
}

// iota 枚举
// iota 在 const 关键字出现时将被重置为 0（const 中每新增一行常量声明将使 iota 增加 1）。
// 如果 iota 定义在 const 定义组中的第 n 行，那么 iota 的值为 n - 1。所以一定要注意 iota 出现在定义组中的第几行，而不是当前代码中它第几次出现。
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func main6() {
	gender := Gender("Female")
	fmt.Println(gender.IsMale())
	fmt.Println(gender.String())

	fmt.Println(January)

}

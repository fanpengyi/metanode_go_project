package basecode

import (
	"fmt"
	"strconv"
)

func main_str_turn() {
	//var i int32 = 17
	//var b byte = 5
	//var f float32
	//
	//// 数字类型可以直接强转
	//f = float32(i) / float32(b)
	//fmt.Printf("f 的值为: %f\n", f)
	//
	//// 当int32类型强转成byte时，高位被直接舍弃
	//var i2 int32 = 256
	//var b2 byte = byte(i2)
	//fmt.Printf("b2 的值为: %d\n", b2)

	//str := "hello,123"
	//bytes := []byte(str)
	//runes := []rune(str)
	//fmt.Printf("bytes 的值: %v\n", bytes)
	//fmt.Printf("runes 的值: %v\n", runes)
	//
	//str2 := string(bytes)
	//str3 := string(runes)
	//
	//fmt.Printf("str2 的值: %s\n", str2)
	//fmt.Printf("str3 的值: %s\n", str3)
	//
	//
	//strconv 可以把数字转成字符串，也可以把字符串转换成数字。

	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err) // 如果转换失败，抛出异常
	} // 将字符串转换为整数
	fmt.Printf("num 的值为: %d\n", num)
	str1 := strconv.Itoa(num) // 将整数转换为字符串
	fmt.Printf("str1 的值为: %s\n", str1)

	// strconv 包还提供了 ParseInt、ParseFloat、FormatInt、FormatFloat 等函数来处理不同类型的转换。
	parseUint, err := strconv.ParseUint(str, 10, 32)
	fmt.Printf("parseUint 的值为: %d\n", parseUint)

	formatUint := strconv.FormatUint(parseUint, 2)
	fmt.Printf("formatUint 的值为: %s\n", formatUint)

}

func main_type_turn() {

	//var i interface{} = "3b"
	//a, ok := i.(int)
	//if ok {
	//	fmt.Println("a is int:", a)
	//} else {
	//	fmt.Println("a is not int")
	//}

	var i interface{} = "ab3123"

	switch i.(type) {
	case int:
		fmt.Println("i is int")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is unknown type")
	}

}

// 定义一个接口
type Supplier interface {
	Get() string
}

//定义一个结构体实现接口

type DigitSupplier struct {
	value int
}

func (d *DigitSupplier) Get() string {
	sprintf := fmt.Sprintf("DigitSupplier value is %d", d.value)
	return sprintf
}

func main_turn_interface() {
	//声明一个接口，指向一个结构体 使用指针
	var supplier Supplier = &DigitSupplier{11}
	fmt.Println(supplier.Get())
	// 类型断言，判断接口类型是否为 DigitSupplier
	var b, ok = (supplier).(*DigitSupplier)
	fmt.Printf("b is %v, ok is %v\n", b, ok)

}

type SameFieldA struct {
	name  string
	value int
}

type SameFieldB struct {
	name  string
	value int
}

func (s *SameFieldB) getValue() int {
	return s.value
}

func main16() {
	a := SameFieldA{
		name:  "a",
		value: 1,
	}

	b := SameFieldB(a)
	fmt.Printf("conver SameFieldA to SameFieldB, value is : %d \n", b.getValue())

	// 只能结构体类型实例之间相互转换，指针不可以相互转换
	// var c interface{} = &a
	// _, ok := c.(*SameFieldB)
}

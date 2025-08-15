package main

import "fmt"

// 定义一个支付接口
// 1 接口可以嵌套。
type PayMethod interface {
	Account
	Pay(amount int) bool
}

// 定义一个账户接口
type Account interface {
	GetBalance() int
}

// 定义一个信用卡结构体
type CreditCart struct {
	balance int
	limit   int
}

func (c *CreditCart) Pay(amount int) bool {

	if c.balance+amount <= c.limit {
		c.balance = c.limit - amount
		fmt.Printf("信用卡支付成功 : %d \n", amount)
		return true
	}
	fmt.Println("信用卡支付失败，超过限额")
	return false
}

// 结构体实现方法
func (c *CreditCart) GetBalance() int {
	return c.balance
}

// 定义一个借记卡结构体
type DebitCart struct {
	balance int
}

func (d *DebitCart) Pay(amount int) bool {
	if d.balance >= amount {
		d.balance -= amount
		fmt.Printf("借记卡支付成功 : %d \n", amount)
		return true
	}
	fmt.Println("借记卡支付失败，余额不足")
	return false
}

func (d *DebitCart) GetBalance() int {
	return d.balance
}

//使用PaymentMethod 接口函数

func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买商品成功，支付金额：%d，当前余额：%d\n", price, p.GetBalance())
	} else {
		fmt.Printf("购买商品失败，支付金额：%d，当前余额：%d\n", price, p.GetBalance())
	}
}

// Go 中接口声明的方法并不要求需要全部公开。
// 如果方法都是值接收者，可以用值或指针。[如果任何方法是指针接收者，则必须用指针。]
func main_interface() {

	creditCart := &CreditCart{0, 1000}

	debitCart := &DebitCart{500}

	fmt.Println("使用信用卡支付")
	purchaseItem(creditCart, 800)

	fmt.Println("使用借记卡支付")
	purchaseItem(debitCart, 300)

	fmt.Println("再次使用借记卡支付")
	purchaseItem(debitCart, 300)

}

// 接口中声明的方法，参数可以没有名称。
type PayMethod2 interface {
	Pay(int)
}

type CreditCart2 struct {
	balance int
	limit   int
}

func (c *CreditCart2) Pay(amount int) {
	if c.balance < amount {
		fmt.Printf("信用卡支付失败，余额不足，当前余额：%d \n", c.balance)
	}
	c.balance -= amount
}

// 如果函数参数使用 interface{}可以接受任何类型的实参。同样，可以接收任何类型的值也可以赋值给 interface{}类型的变量
func anyParam(param interface{}) {
	fmt.Printf("param is %v, type is %T\n", param, param)
}

func main17() {

	cart2 := CreditCart2{1000, 2000}
	cart2.Pay(200)
	var p PayMethod2 = &cart2
	fmt.Println("p.Pay:", p)

	var b interface{} = &cart2
	fmt.Println("b is interface:", b)

	anyParam(cart2)
	anyParam(1)
	anyParam("123")
	anyParam(p)

}

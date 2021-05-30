package main

import "fmt"

type U struct {
}
type Pay interface {
	Order(id int) int
}

type WeChat struct {
	//pay string
}
type AlPay struct {
	pay string
}

func (w *WeChat) Order(id int) int {
	return 200
}

func (a *AlPay) Order(id int) int {
	return 300
}

func main()  {
	//a :=1
	//b := &a
	//fmt.Printf("a:%d ptr: %p t:%T \n", a, &a, a)
	//fmt.Printf("b:%p ptr: %p t:%T \n", b, &b, b)

	var p  Pay
	p = &WeChat{}
	fmt.Println(p.Order(3))

	var t interface{}
	t = &WeChat{}
	switch t.(type) {
	case int:
		fmt.Println("int")
	case *WeChat:
		fmt.Println("wechat")

	}
}

func NewU() *U {
	return &U{}
}

func arr(r int) *int {
	r++
	return &r
}

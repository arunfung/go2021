package main

import "fmt"

type Pay interface {
	Order(id int) int
	Syn()
}

type WeChat struct {
	pay string
}
type AlPay struct {
	pay string
}

func (w *WeChat) Syn() {
}

func (w *WeChat) Order(id int) int {
	return 200
}

func (a *AlPay) Order(id int) int {
	return 300
}

func p(pay Pay, id int) {
	fmt.Println(pay.Order(id))
}

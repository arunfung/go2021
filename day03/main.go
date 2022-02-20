package main

import (
	"errors"
	"fmt"
)

func main() {
	//var arr = index
	//arr()
	//fmt.Println(arr)

	//sum(1,2,3,4)
	//
	//test := func(k int) {
	//	fmt.Println("这是闭包", k)
	//}
	//test(1)
	//tests := tests()
	//tests()
	//tests()
	//tests()
	//fmt.Println(def())
	//pa()

	//defer func() {
	//	fmt.Println("捕获异常", recover())
	//}()
	//panic("运行异常")

	//try(func() {
	//	fmt.Println("运行")
	//	panic("运行异常")
	//}, func(err interface{}) {
	//	fmt.Println("捕获的异常", err)
	//})
	i, err := sub()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}
func sub() (int, error) {
	return 9, errors.New("这是new 的异常")
}

func try(fun func(), catch func(err interface{})) {
	defer func() {
		fmt.Println("捕获异常，处理中")
		if err := recover(); err != nil {
			fmt.Println("捕获异常")
			catch(err)
		}
	}()
	fun()
}

func tests() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}

func index() {
	fmt.Println("ddd")
}

func sum(a ...int) {
	fmt.Println(a)
}

func def() int {
	// 延迟执行
	defer index()
	sum(4, 5, 6)
	sum(1, 2, 3)
	fmt.Println("这是 def")
	return 9
}

func pa() {
	panic("这是异常")
}

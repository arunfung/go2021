package controller

import (
	"arunfung/go2021/edu/util"
	"fmt"
	"strconv"
)

var autoController *AutoController
func init() {
	autoController = &AutoController{}
}

func Run() {

	dispatch()

}

func dispatch() {

	fmt.Println("<<<======= start ======>>> ")
	fmt.Println("你要执行的操作")
	oper := []string{"登入", "注册"}
	for k, v := range oper {
		fmt.Printf("(%d) : %s\n", k, v)
	}

	flag, _ := strconv.Atoi(util.CRead())
	fmt.Println("<<<======= end ======>>> ")

	fmt.Println("<<<======= start ======>>> ")
	// if 是否为登入 =>
	switch flag {
	case 0:
		al := autoController.login()
		if al {
			fmt.Println("<<<======= 登陆成功 ======>>> ")
		}
	case 1:
	}

}
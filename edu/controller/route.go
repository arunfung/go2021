package controller

import (
	"arunfung/go2021/edu/util"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	next string
	view string
)

func Run() {

	// 初次访问的方法
	next = "index::Welcome"


	for {
		flag := util.CReturn(util.Cfunc(dispatch))
		if flag {
			break
		}
	}

	fmt.Println("结束")

}

func dispatch() (bool, error){
	// 1. 获取展示的列表
	args := strings.Split(next, "::")
	controller, ok := controllers[args[0]]
	if ok != true {
		return false, errors.New("无法根据指定标识查找控制器: " + args[0])
	}
	// 2. 执行方法
	// 在控制器中会设置view也就是下次视图会展示的内容
	reflect.ValueOf(controller).MethodByName(args[1]).Call([]reflect.Value{})

	// 3. 获取下一步的运算
	opers, ok := views[view]
	if ok != true {
		return false, errors.New("无法根据指定标识查找到视图: " + args[0])
	}
	// fmt.Println(opers)
	// 4. 处理输出格式
	fmethods, fdesc := toModelFormat(opers)
	// 输出下一步操作信息
	util.Coper(fdesc)

	// 5. 匹配操作转发下次执行
	for {
		input := util.CRead()
		if input == "x" {
			return true, nil
		}
		flag, err := strconv.Atoi(input)
		if err == nil && flag < len(fmethods) {
			next = fmethods[flag]
			break
		}
		fmt.Println("信息输入有误请重新输入")
	}
	return false, nil
}
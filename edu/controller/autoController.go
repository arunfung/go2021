package controller

import (
	"arunfung/go2021/edu/model"
	"arunfung/go2021/edu/util"
	"fmt"
)

type AutoController struct{}

func (c *AutoController) Login() {
	fmt.Println("输入用户名")
	username := util.CRead()
	fmt.Println("输入密码")
	password := util.CRead()
	user := model.GetUser(username)

	if user == nil {
		fmt.Println("查询不到用户", username)
		return
	}
	if user.GetPassword() == password {
		fmt.Println("登入成功")
		view = "index_view"
		return
	} else {
		fmt.Println("密码错误")
		view = "login_view"
		return
	}
}

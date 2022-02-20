package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"hello/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	name := c.GetString("name")
	user, err := models.FindUserByName(name)
	fmt.Println("user", user)
	if err != nil {
	}

	c.SetSession("uid", user.Id)
	res := make(map[string]interface{})
	res["code"] = 200
	res["data"] = user
	c.Data["json"] = res
	c.ServeJSON()
}

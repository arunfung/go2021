package main

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_ "hello/initial"
	_ "hello/routers"
)

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Run()
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	if !ok && ctx.Request.RequestURI != "/login" {
		//ctx.Redirect(302, "/login")
		fmt.Println("跳转")
	}
}

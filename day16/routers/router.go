package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
}

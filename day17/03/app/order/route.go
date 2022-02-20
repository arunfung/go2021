package order

import (
	"arunfung/go2021/day17/03/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	routes.RegisterRoute(Routes)
}

func Routes(g *gin.Engine) {
	g.GET("/getorder", GetOrder)
}

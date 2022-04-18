package api

import (
	apiV1 "client/app/api/v1"
	"client/routes"

	"github.com/gin-gonic/gin"
)

// 初始化的时候注册
func init() {
	routes.Register(Routes)
}

func Routes(g *gin.Engine) {

	goods := g.Group("goods")
	{
		goods.GET("", apiV1.GetGoods)
	}
}

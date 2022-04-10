package api

import (
	apiV1 "frame/app/api/v1"
	"frame/routes"
	"github.com/gin-gonic/gin"
)

// 初始化的时候注册
func init() {
	routes.Register(Routes)
}

func Routes(g *gin.Engine) {

	goods := g.Group("goods")
	{
		goods.POST("", apiV1.GetGoods)
	}
}

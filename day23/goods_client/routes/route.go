package routes

import (
	"frame/middleware"
	"github.com/gin-gonic/gin"
)

type Router func(*gin.Engine)

var routers = []Router{}

// 初始化gin路由
func InitRoutes() *gin.Engine {
	r := gin.Default()

	// 注册中间件
	middleware.InitMiddleware(r)

	for _, route := range routers {
		route(r)
	}
	return r
}

// 注册route
func Register(routes ...Router) {
	routers = append(routers, routes...)
}

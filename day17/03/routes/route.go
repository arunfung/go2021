package routes

import "github.com/gin-gonic/gin"

type Router func(engine *gin.Engine)

var routers = []Router{}

func RegisterRoute(routes ...Router)  {
	routers = append(routers, routes...)
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	for _, route := range routers {
		route(r)
	}
	return r
}
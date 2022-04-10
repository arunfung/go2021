package main

import (
	_ "frame/app"
	"frame/routes"
	"go-micro.dev/v4/web"
)

func main() {
	gin := routes.InitRoutes()
	server := web.NewService(
		web.Name("goods.client"),
		web.Handler(gin),
	)

	server.Init()
	gin.Run()
}

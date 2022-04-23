package main

import (
	_ "client/app"
	"client/routes"
	"client/rpc"

	"go-micro.dev/v4/web"
)

func main() {
	gin := routes.InitRoutes()
	server := web.NewService(
		web.Name("goods.client"),
		web.Handler(gin),
		web.Registry(rpc.ConsulReg),
	)

	server.Init()
	gin.Run()
}

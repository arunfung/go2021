package main

import (
	_ "frame/app"
	"frame/routes"
)

func main() {
	gin := routes.InitRoutes()
	gin.Run()
}

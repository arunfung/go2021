package main

import (
	_ "arunfung/go2021/day17/03/app"
	"arunfung/go2021/day17/03/routes"
)

func main() {
	routes.InitRouter().Run()
}

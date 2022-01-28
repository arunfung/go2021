package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(InitApp())
	r.Use(Loginc())
	r.GET("/index", func(c *gin.Context) {
		fmt.Println("==>>> index  ====>>")

		c.JSON(200, gin.H{
			"code": "ok",
		})
	})

	r.GET("/user", Auth(), func(c *gin.Context) {
		fmt.Println("==>>> user  ====>>")
		c.JSON(200, gin.H{
			"msg": "user",
		})
	})
	// 默认端口是 8080
	r.Run()
}

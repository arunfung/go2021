package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": "ok"})
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/login/:name", login)
	}
	r.POST("/login", login2)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func login(c *gin.Context) {
	name := c.Param("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"code": "ok",
		"name": name,
		"age":  age,
	})
}
func login2(c *gin.Context) {
	c.String(200, "hello gin post")
}

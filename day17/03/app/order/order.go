package order

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetOrder(c *gin.Context) {
	fmt.Println("get order ===>>> start")
	c.String(200, "get order")
}

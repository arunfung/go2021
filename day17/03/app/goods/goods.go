package goods

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetGoods(c *gin.Context) {
	fmt.Println("get goods ===>>> start")
	c.String(200, "get goods")
}

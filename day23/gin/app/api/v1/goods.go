package v1

import (
	"frame/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGoods(c *gin.Context) {
	goodsId := c.Query("goods_id")
	goodsInfo := services.GetGoodsDetail(map[string]interface{}{
		"id": goodsId,
	})
	c.JSON(http.StatusOK, goodsInfo)
}

package v1

import (
	"frame/rpc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetGoods(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id"))

	goodsInfo := rpc.GetGoodsDetail(goodsId)
	c.JSON(http.StatusOK, goodsInfo)
}

package v1

import (
	"client/rpc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WarpGoods(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id"))

	goodsInfo, err := rpc.GetGoodsDetail(goodsId) // micro

	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, goodsInfo)
	}
}

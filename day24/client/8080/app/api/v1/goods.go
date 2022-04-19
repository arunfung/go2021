package v1

import (
	"client/rpc"
	"net/http"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func GetGoods(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id"))

	// 1. 定义hystrix配置
	configHy := hystrix.CommandConfig{
		Timeout: 1000, // 单位毫秒
	}
	// 2. 配置commad
	hystrix.ConfigureCommand("goodsHy", configHy)
	// 3. 使用hystrix
	var goodsInfo interface{}
	var err error

	err = hystrix.Do("goodsHy", func() error {
		goodsInfo, err = rpc.GetGoodsDetail(goodsId)
		return err
	}, func(err error) error {
		// 降级从缓存读取
		goodsInfo = Demotion(goodsId)
		return nil
	})

	// err 的异常情况 hystrix：熔断、 程序调用rpc出现服务异常
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, goodsInfo)
	}
}

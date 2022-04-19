package v1

import (
	proto "client/proto"
	"client/rpc"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var goodsKey = "goods_redis_%d"

func Wredis(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id"))
	goodsInfo, _ := rpc.GetGoodsDetail(goodsId)
	goodsInfo.Msg = "write redis"

	// 连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	data, _ := json.Marshal(goodsInfo)

	key := fmt.Sprintf(goodsKey, goodsId)
	// 写入数据
	rdb.Set(c, key, data, 0)
	c.JSON(http.StatusOK, goodsInfo)
}

func Gredis(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id"))

	// 连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	key := fmt.Sprintf(goodsKey, goodsId)
	// 写入数据
	data, _ := rdb.Get(c, key).Result()
	goodsInfo := &proto.ResponseGoods{}

	_ = json.Unmarshal([]byte(data), goodsInfo)

	c.JSON(http.StatusOK, goodsInfo)
}

func Demotion(id int) *proto.ResponseGoods {
	// 连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	key := fmt.Sprintf(goodsKey, id)
	// 写入数据
	data, _ := rdb.Get(context.Background(), key).Result()
	goodsInfo := &proto.ResponseGoods{}

	_ = json.Unmarshal([]byte(data), goodsInfo)
	return goodsInfo
}

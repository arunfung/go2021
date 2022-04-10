package rpc

import (
	"context"
	"fmt"
	proto "frame/proto"
	"go-micro.dev/v4/client"
)

func GetGoodsDetail(goodsId int) *proto.ResponseGoods {
	// 创建客户端
	goods := proto.NewGoodsService("goods", client.DefaultClient)
	fmt.Println("GetGoodsDetail")
	res, err := goods.GetGoodsDetail(context.Background(), &proto.RequestGoods{
		GoodsId: int32(goodsId),
	})
	fmt.Println("err : ", err)
	return res
}

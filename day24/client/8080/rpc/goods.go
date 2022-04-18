package rpc

import (
	proto "client/proto"
	"context"
	"fmt"

	"go-micro.dev/v4/client"
)

func GetGoodsDetail(goodsId int) (*proto.ResponseGoods, error) {
	// 创建客户端
	goods := proto.NewGoodsService("goods", client.DefaultClient)
	fmt.Println("GetGoodsDetail")
	res, err := goods.GetGoodsDetail(context.Background(), &proto.RequestGoods{
		GoodsId: int32(goodsId),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

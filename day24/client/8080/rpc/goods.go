package rpc

import (
	proto "client/proto"
	"client/wrapper/hystrix"
	"client/wrapper/login"
	"context"
	"fmt"

	"go-micro.dev/v4"
)

var goodsCli proto.GoodsService

func init() {
	hystrix.ConfigureDefault(hystrix.CommandConfig{Timeout: 4000})
	microClient := micro.NewService(
		micro.Name("goods.client"),
		micro.WrapClient(hystrix.NewClientWrapper()),
		micro.WrapClient(login.NewClientWrapper()),
	)
	// 创建客户端
	goodsCli = proto.NewGoodsService("goods", microClient.Client())
}

func GetGoodsDetail(goodsId int) (*proto.ResponseGoods, error) {

	fmt.Println("GetGoodsDetail")
	res, err := goodsCli.GetGoodsDetail(context.Background(), &proto.RequestGoods{
		GoodsId: int32(goodsId),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

package rpc

import (
	proto "client/proto"
	"client/wrapper/hystrix"
	"client/wrapper/login"
	"context"
	"fmt"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

var ConsulReg registry.Registry
var goodsCli proto.GoodsService

func init() {
	// 注册客户端服务
	ConsulReg = consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	hystrix.ConfigureDefault(hystrix.CommandConfig{Timeout: 1000})
	ops := hystrix.WithFallback(func(ctx context.Context, err error) error {

		return err
	})
	microClient := micro.NewService(
		micro.Name("goods.client"),
		micro.WrapClient(hystrix.NewClientWrapper(ops)),
		micro.WrapClient(login.NewClientWrapper()),
	)
	// 创建RPC客户端
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

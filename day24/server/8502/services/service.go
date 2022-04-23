package services

import (
	proto "server/proto"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

var consulReg registry.Registry

func init() {
	consulReg = consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
}

func Run() {
	// 创建服务
	service := micro.NewService(
		micro.Name("goods"),
		micro.Address(":8502"),
		micro.Registry(consulReg),
	)
	//初始化服务
	service.Init()
	// 注册服务
	proto.RegisterGoodsServiceHandler(service.Server(), new(RequestGoods))

	service.Run()
}

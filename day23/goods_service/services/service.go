package services

import (
	proto "frame/proto"
	"go-micro.dev/v4"
)

func Run() {
	// 创建服务
	service := micro.NewService(
		micro.Name("goods"),
		//micro.Address()
	)
	//初始化服务
	service.Init()
	// 注册服务
	proto.RegisterGoodsServiceHandler(service.Server(), new(RequestGoods))

	service.Run()
}

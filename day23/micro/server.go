package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	proto "micro/proto"
)

func main() {
	// 创建服务
	service := micro.NewService(
		micro.Name("greeter.service"),
		micro.Address(":9630"),
	)
	//初始化服务
	service.Init()

	// 注册服务
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	service.Run()
}

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, res *proto.HelloResponse) error {
	fmt.Println("servier", req.Name)
	res.Greeting = "show : " + req.Name
	return nil
}

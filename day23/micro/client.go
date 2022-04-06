package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	proto "micro/proto"
)

func main() {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// 创建客户端
	greeter := proto.NewGreeterService("greeter.service", service.Client())
	// 调用方法
	res, err := greeter.Hello(context.Background(), &proto.HelloRequest{
		Name: "arunfung",
	})
	fmt.Println("res : ", res)
	fmt.Println("err ； ", err)
}

package main

import (
	pb "day21/proto/goods" // 引入编译生成的包
	"fmt"
	"net" // 引入编译生成的包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	ADDRESS = "127.0.0.1:1234"
)

type GoodsService struct {
	pb.UnimplementedGoodsServiceServer
}

var GoodsServices = GoodsService{}

func (g GoodsService) Get(ctx context.Context, in *pb.Goods) (*pb.Goods, error) {
	goods := &pb.Goods{
		Id:   1,
		Name: "苹果",
	}
	return goods, nil
}

func main() {
	listen, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		fmt.Println("err : ", err)
	}

	//creds, _ := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	// 实例化grpc server
	//s := grpc.NewServer(grpc.Creds(creds))
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterGoodsServiceServer(s, GoodsServices)

	// grpc监听端口和地址
	fmt.Println("listen on : " + ADDRESS)
	s.Serve(listen)
}

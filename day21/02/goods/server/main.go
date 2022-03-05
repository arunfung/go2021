package main

import (
	pb "day21/proto/goods" // 引入编译生成的包
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net" // 引入编译生成的包
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

	cer := "../../keys/server.pem"
	key := "../../keys/server.key"
	creds, _ := credentials.NewServerTLSFromFile(cer, key)
	// 实例化grpc server
	s := grpc.NewServer(grpc.Creds(creds))
	//s := grpc.NewServer()
	// 注册服务
	pb.RegisterGoodsServiceServer(s, &GoodsService{})

	// grpc监听端口和地址
	fmt.Println("listen on : " + ADDRESS)
	s.Serve(listen)
}

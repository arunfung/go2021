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

const Cer = "/Users/arun/go/src/arunfung/go2021/day21/03/keys/server.pem"
const Key = "/Users/arun/go/src/arunfung/go2021/day21/03/keys/server.key"

type GoodsService struct {
	pb.UnimplementedGoodsServiceServer
}

var GoodsServices = GoodsService{}

func (g GoodsService) Get(ctx context.Context, in *pb.Goods) (*pb.Goods, error) {
	//Auth(ctx)
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
	var opts []grpc.ServerOption
	creds, _ := credentials.NewServerTLSFromFile(Cer, Key)
	opts = append(opts, grpc.Creds(creds))
	opts = append(opts, grpc.UnaryInterceptor(intercaptor))
	// 实例化grpc server
	s := grpc.NewServer(opts...)
	//s := grpc.NewServer()
	// 注册服务
	pb.RegisterGoodsServiceServer(s, &GoodsService{})

	// grpc监听端口和地址
	fmt.Println("listen on : " + ADDRESS)
	s.Serve(listen)
}

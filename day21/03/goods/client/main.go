package main

import (
	pb "day21/proto/goods" // 引入编译生成的包
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:1234"
	OpenSSL = true
)

const Cer = "/Users/arun/go/src/arunfung/go2021/day21/03/keys/server.pem"

// 自定义验证器
type customCreds struct{}

// GetRequestMetadata 实现自定义认证接口
func (c customCreds) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"userId": "1",
		"token":  "token",
	}, nil
}

// RequireTransportSecurity 验证是否开启ssl
func (c customCreds) RequireTransportSecurity() bool {
	return OpenSSL
}

func main() {
	var err error
	var opts []grpc.DialOption

	if OpenSSL {
		// 开启证书的验证方式
		creds, _ := credentials.NewClientTLSFromFile(Cer, "")

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		// 正常的验证
		opts = append(opts, grpc.WithInsecure())
	}
	// 使用自定义的验证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCreds)))
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	// 连接
	conn, err := grpc.Dial(Address, opts...)
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewGoodsServiceClient(conn)

	// 调用方法
	req := &pb.Goods{}
	res, err := c.Get(context.Background(), req)

	if err != nil {
		grpclog.Fatalln(err)
	}

	// grpclog.Println(res.Message)
	fmt.Println(res)
}

// interceptor 客户端拦截器
func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Println("client : 拦截器")
	return err
}

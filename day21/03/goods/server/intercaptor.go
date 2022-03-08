package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func intercaptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	Auth(ctx)
	fmt.Println("进入拦截器")
	return handler(ctx, req)
}

package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
)

const (
	TOKEN = "token" // jwt
)

func Auth(ctx context.Context) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("无token信息")
		return false, errors.New("无token信息")
	}

	userId := md["userId"]
	token := md["token"]
	fmt.Println("userId : ", userId, " token : ", token)

	return true, nil
}

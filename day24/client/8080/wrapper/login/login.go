package login

import (
	"context"
	"fmt"

	"go-micro.dev/v4/client"
)

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("====>>> wrapper login ===>> start")
	// 业务处理逻辑
	err := c.Client.Call(ctx, req, rsp, opts...)
	fmt.Println("====>>> wrapper login ===>> end")
	return err
}

// NewClientWrapper returns a hystrix client Wrapper.
func NewClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}

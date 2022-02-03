package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Goods struct {
	Id   int
	Name string
}

func (g *Goods) FindById(args *Goods, reply *Goods) error {
	fmt.Println("参数:", args)
	*reply = *g
	return nil
}
func (g *Goods) FindGoodNameById(args *Goods, reply *string) error {
	fmt.Println("参数:", args)
	*reply = g.Name
	return nil
}

func main() {
	goods := &Goods{
		Id:   1,
		Name: "arun",
	}
	// 注册对外服务
	rpc.Register(goods)
	rpc.HandleHTTP()
	// 开启rpc监听 -》端口和ip
	listen, _ := net.Listen("tcp", ":1234")
	defer listen.Close()
	go http.Serve(listen, nil)
	time.Sleep(100e9)
}

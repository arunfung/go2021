package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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

	// 开启rpc监听 -》端口和ip
	listen, _ := net.Listen("tcp", ":1234")
	defer listen.Close()
	// 建立连接
	conn, _ := listen.Accept()
	rpc.RegisterName("Goods", goods)
	jsonrpc.ServeConn(conn)
}

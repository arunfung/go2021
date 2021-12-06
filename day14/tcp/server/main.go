package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("启动服务端 ： tcp://127.0.0.1:3333")
	// 1. 监听端口 tcp://0.0.0.0:3333  监听的网络主要以本机可用ip为主
	listen, err := net.Listen("tcp", "127.0.0.1:3333")
	if err != nil {
		fmt.Println("err : ", err)
		return 
	}

}

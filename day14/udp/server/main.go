package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("启动服务端监听 ： udp://127.0.0.1:5555")
	// 1. 监听地址 传递net.UDPAddr struct
	listen, _ := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 5555,
	})
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		fmt.Println("data : ", string(data[:n]), addr)
		if err != nil {
			fmt.Println("err : ", err)
			continue
		}
		listen.WriteToUDP([]byte("im is udp"), addr)
	}
}

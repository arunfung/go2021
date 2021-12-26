package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建连接
	socket, _ := net.DialUDP("udp", nil,&net.UDPAddr{
		IP: net.IPv4(127, 0, 0, 1),
		Port: 5555,
	})
	fmt.Println("与udp://127.0.0.1:5555 建立连接")
	defer socket.Close()
	// 2.进行数据的发送接收
	socket.Write([]byte("hello udp sever"))
	var data [1024]byte
	n,addr,_ := socket.ReadFromUDP(data[:])
	fmt.Println("n :", string(data[:n]), addr)
}

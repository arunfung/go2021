package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建连接
	conn, _ := net.Dial("tcp", "127.0.0.1:3333")
	fmt.Println("与tcp://127.0.0.1:3333建立连接")
	defer conn.Close()
	// 2.进行数据的发送接收
	conn.Write([]byte("hello sever"))
	var data [1024]byte
	n, _ := conn.Read(data[:])
	fmt.Println("n :", string(data[:n]))
}

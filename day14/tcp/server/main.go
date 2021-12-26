package main

import (
	"bufio"
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
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("err : ", err)
			return
		}
		go handler(conn)
	}
}

func handler(c net.Conn)  {
	defer c.Close()
	for {
		var data [1024]byte
		n, err := bufio.NewReader(c).Read(data[:])
		if err != nil {
			fmt.Println("err : ", err)
			break
		}
		fmt.Println("n", string(data[:n]))

		c.Write([]byte("hello im is server"))
	}
}

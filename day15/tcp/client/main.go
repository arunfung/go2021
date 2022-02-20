package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	// 创建连接
	conn, _ := net.Dial("tcp", "127.0.0.1:3333")
	fmt.Println("与tcp://127.0.0.1:3333建立连接")
	defer conn.Close()

	// int16 2 个字节
	msg := "arunfung"
	msgLen := len(msg)
	length := int16(msgLen)
	fmt.Println("msgLen :", msgLen)
	fmt.Println("length :", length)

	pkg := new(bytes.Buffer)
	binary.Write(pkg, binary.BigEndian, length)

	data := append(pkg.Bytes(), []byte(msg)...)
	fmt.Println("data :", string(data))

	// 2.进行数据的发送接收
	conn.Write(data)
	var pack [1024]byte
	n, _ := conn.Read(pack[:])
	fmt.Println("n :", string(pack[:n]))
}

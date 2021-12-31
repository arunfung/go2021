package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
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

func handler(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		msg, err := uppack(reader)

		if err != nil {
			fmt.Println("err : ", err)
			break
		}
		fmt.Println("n", msg)

		c.Write([]byte("hello im is server"))
	}
}

func uppack(reader *bufio.Reader) (string, error) {
	lenByte, _ := reader.Peek(2)

	lengthBuff := bytes.NewBuffer(lenByte)

	var length int16
	err := binary.Read(lengthBuff, binary.BigEndian, &length)
	if err != nil {
		return "", err
	}

	if int16(reader.Buffered()) < length {
		return "", errors.New("数据异常")
	}

	pack := make([]byte, int(2+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}

	return string(pack[2:]), nil
}

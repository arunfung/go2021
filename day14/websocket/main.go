package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	fmt.Println("启动websocket://:7777")
	// 设定websocket的服务信息处理
	http.Handle("/", websocket.Handler(server))
	// 设定监听
	http.ListenAndServe(":7777", nil)
}

func server(ws *websocket.Conn) {
	fmt.Println("new connection")
	data := make([]byte, 1024)
	for {
		// 读取信息
		d, err := ws.Read(data)
		if err != nil {
			fmt.Println("err : ", err)
			break
		}
		fmt.Println("读取到的信息 ", d)
		ws.Write([]byte("你好I'm is webscoket server"))
	}
}

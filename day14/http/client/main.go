package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://127.0.0.1:8888/bar")
	//fmt.Println("resp", resp)

	var data [1024]byte // [:]  data := make([]byte, 1024)
	n, _ := resp.Body.Read(data[:])
	fmt.Println("信息 : ", string(data[:n]))
}

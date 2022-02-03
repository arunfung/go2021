package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

type Goods struct {
	Id   int
	Name string
}

func main() {
	conn, _ := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	defer conn.Close()
	reqGoods := Goods{Id: 1}
	resGoods := &Goods{}
	err := conn.Call("Goods.FindById", reqGoods, resGoods)
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Println("resGoods : ", resGoods)

	var data string
	err = conn.Call("Goods.FindGoodNameById", reqGoods, &data)
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Println("data : ", data)
}

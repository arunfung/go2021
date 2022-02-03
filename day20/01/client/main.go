package main

import (
	"fmt"
	"net/rpc"
)

type Goods struct {
	Id   int
	Name string
}

func main() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	reqGoods := Goods{Id: 1}
	resGoods := &Goods{}
	err := client.Call("Goods.FindById", reqGoods, resGoods)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("resGoods", resGoods)
	goodName := ""
	err1 := client.Call("Goods.FindGoodNameById", reqGoods, &goodName)
	if err1 != nil {
		fmt.Println("err1:", err1)
	}
	fmt.Println("goodName", goodName)
}

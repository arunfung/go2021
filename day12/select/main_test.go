package main

import (
	"fmt"
	"testing"
	"time"
)


func TestSelect(t *testing.T) {
	goodCh := make(chan string)
	orderCh := make(chan string)

	go goods(goodCh)
	go order(orderCh)
	AAA:
	for {
		select {
		case order := <-orderCh:
			fmt.Println("order", order)
			//return
			break AAA
		case good := <-goodCh:
			fmt.Println("good", good)
		case <-time.After(2e9):
			fmt.Println("超时")
			break AAA
		//default:
		//	fmt.Println("default")
		}
	}
}

func TestTimer(t *testing.T) {
	ticker := time.Tick(1e9)
	for {
		<- ticker
		go curl()
	}
}
func curl() {
	fmt.Println("请求某一个接口")
}
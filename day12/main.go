package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(4) // 设置逻辑CPU的个数
	ch := make(chan int)  // 缓冲区
	wg.Add(1)
	go send(ch)
	wg.Wait()
}

func send(ch chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go recv(ch)
		ch <- i
	}
}

func recv(ch chan int) {
	defer wg.Done()
	fmt.Println("接收通道信息", <-ch) // 通道没有数据，阻塞
}
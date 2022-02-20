package main

import (
	"fmt"
)

// 阻塞是否影响 main 运行
func main() {
	ch := make(chan int)
	done := make(chan bool)
	go send(0, 10, ch) // 执行 1-10 结束，如果执行完不关闭 ch 通道
	go recv(ch, done)  // 循环从通道获取信息阻塞-》导致 done 通道无法写入数据
	<-done             // done 通道就无法获取数据，一直阻塞就会影响 main 运行
}

func send(start, end int, ch chan<- int) {
	for i := start; i < end; i++ {
		ch <- i
	}
	close(ch)
}

func recv(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Println("接收通道信息", num)
	}
	done <- true
}

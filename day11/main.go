package main

import (
	"fmt"
	"runtime"
	"time"
)

var sendChOnly chan<- string // 单向发送通道
var recvChOnly <-chan string // 单向获取通道

func main() {
  	//runtime.GOMAXPROCS(1)
	//go run1()
	//go run2()
	//time.Sleep(time.Second)

	//ch := make(chan string)
	//go send(ch)
	//go receive(ch)
	//time.Sleep(time.Second)

	//ch := make(chan int, 1)
	//ch <- 10
	//x := <-ch
	//xs := <-ch
	//fmt.Println("x", x)
	//fmt.Println("x", xs)

	//ch := make(chan string, 1)
	//sendChOnly = ch
	//sendChOnly <- "arun"
	//recvChOnly = ch
	//fmt.Println(<-recvChOnly)

	ch := make(chan string, 1)
	go sendfChOnly(ch)
	go recvfChOnly(ch)
	time.Sleep(time.Second)
}

func run1()  {
	fmt.Println("start run1")
	runtime.Gosched()
	fmt.Println("end run1")
}

func run2()  {
	fmt.Println("start run2")
	//runtime.Goexit()
	runtime.Gosched()
	fmt.Println("end run2")
}

func send(ch chan string)  {
	ch <- "arun"
	fmt.Println("1")
	ch <- "arun fung"
	fmt.Println("2")
	ch <- "go"
	fmt.Println("3")
}

func receive(ch chan string)  {
	s := <-ch
	fmt.Println("s : ", s)
	fmt.Println("接收通道信息", <-ch) // 阻塞、切换协程

	if <-ch == "go" {
		fmt.Println("go")
	}
}

func sendfChOnly(ch chan<- string) {
	ch <- "arun fung"
}

func recvfChOnly(ch <-chan string) {
	fmt.Println("接收通道信息", <-ch)
}
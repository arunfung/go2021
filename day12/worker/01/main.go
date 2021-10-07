package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main() {
	ch := make(chan string)
	done := make(chan []string)
	data := make([]string, 0)

	wg.Add(1)
	go redis(ch)
	go func(data []string, ch chan string, done chan []string) {
		for {
			if d, ok := <-ch; ok { // 只有在协程通道关闭的时候才能跳出循环
				data = append(data, d)
			} else {
				break
			}
		}
		fmt.Println("data:", data)
		done <- data
	}(data,ch,done)

	wg.Wait() // 在添加的协程个数，执行完之后就会取消阻塞
	close(ch) // 这是关闭ch通道
	fmt.Println("<-done", <-done) // 需要在 done 通道有信息的时候才执行
}

func redis(ch chan<- string) {
	defer wg.Done()
	ch <- "redis"
}
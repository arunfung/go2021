package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
// 筛选奇数和被三整除的数
func main() {
	ch := make(chan int)
	oddOut := make(chan int)
	threeOut := make(chan int)
	done := make(chan bool)
	go generate(ch, 1, 10, done)
	go filter(ch, oddOut, threeOut)

	for {
		if <- done {
			break
		}
		go func(oddOut chan int) {
			defer wg.Done()
			fmt.Println("奇数", <- oddOut)
		}(oddOut)
		go func(threeOut chan int) {
			defer wg.Done()
			fmt.Println("被三整除", <- threeOut)
		}(threeOut)
	}
	wg.Wait()
}

func generate(ch chan int, start,end int, done chan bool)  {
	for i := start; i <= end; i++ {
		done <- false
		ch <- i
	}
	done <- true
	close(ch)
}

func filter(ch, oddOut, threeOut chan int)  {
	for {
		if i, ok := <-ch; ok {
			if i%2 != 0 {
				wg.Add(1)
				oddOut <- i
			}
			if i%3 == 0 {
				wg.Add(1)
				threeOut <- i
			}
		} else {
			break
		}
	}
}

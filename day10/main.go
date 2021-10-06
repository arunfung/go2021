package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func main() {
	fmt.Println(time.Now())
	wg.Add(1)
	go redis()
	wg.Add(1)
	go mysql()
	wg.Add(1)
	go file()
	wg.Wait()
	fmt.Println(time.Now())
}

func redis() {
	defer wg.Done()
	fmt.Println("start 读取 redis 中的信息")
	time.Sleep(time.Second * 1)
	fmt.Println("end 读取 redis 中的信息")
}

func mysql() {
	defer wg.Done()
	fmt.Println("start 读取 mysql 中的信息")
	time.Sleep(time.Second * 2)
	fmt.Println("end 读取 mysql 中的信息")
}

func file() {
	defer wg.Done()
	fmt.Println("start 读取 file 中的信息")
	time.Sleep(time.Second * 3)
	fmt.Println("end 读取 file 中的信息")
}
package main

import (
	"strconv"
	"sync"
	"testing"
)
var rwlock sync.RWMutex
var count int
var wg sync.WaitGroup

func TestM(t *testing.T) {
	file := "r.txt"
	for i := 0; i < 8000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rwlock.RLock() // 加读锁
			ReadFile(file)
			// fmt.Println("d", d)
			rwlock.RUnlock() // 释放读锁
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rwlock.Lock() // 加读锁
			count++
			// fmt.Println("写入数据", count)
			WriteFile(file, strconv.Itoa(count))
			rwlock.Unlock() // 释放读锁
		}()
	}
	wg.Wait()

}

package main

import (
	"fmt"
	"testing"
)

func TestArr(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p %d\n", &arr1, cap(arr1))

	arr2 := append(arr1, 6)
	fmt.Printf("%p %d\n", &arr2, cap(arr2))
}

func TestMySlice(t *testing.T) {
	MySlice := NewMySlice(2, 2)
	fmt.Println(MySlice.GetLen())
	fmt.Println(MySlice.arr.GetName())
}

func TestIn(t *testing.T) {
	w := &WeChat{}
	p(w, 2)

	//a := &AlPay{}
	//p(a, 2)
}

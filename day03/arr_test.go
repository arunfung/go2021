package main

import (
	"fmt"
	"testing"
)

func TestArr(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p %d\n",&arr1, cap(arr1))

	arr2 := append(arr1, 6)
	fmt.Printf("%p %d\n",&arr2, cap(arr2))
}

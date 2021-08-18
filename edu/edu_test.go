package edu_test

import (
	"fmt"
	"testing"
)

type myFun func() int

func (f myFun) call() int {
	return f()
}

func TestFu(t *testing.T) {
	s := sum
	//fmt.Println(s)
	nuw := myFun(s)
	fmt.Println(nuw.call())
}
func sum() int {
	return 1
}

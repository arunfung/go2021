package calc

import (
	"testing"
)

func TestOper(t *testing.T) {
	num := 1
	num2 := 3
	flag := "+"
	ret := OperationFactory(flag).Exe(num, num2)
	t.Log(ret)
}
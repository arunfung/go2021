package calc

import (
	"testing"
)

func TestOper(t *testing.T) {
	num1 := 1
	num2 := 3
	flag := "+"
	//ret := OperationFactory(flag).Exe(num, num2)
	//t.Log(ret)

	ret := OperationFactory(num1, num2, flag).Exe()
	t.Log(ret)
}

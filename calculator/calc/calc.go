package calc

type OperationInterface interface {
	Exc(a, b int) int
}

type OperationAdd struct {}
type OperationSub struct {}
type OperationMul struct {}
type OperationDiv struct {}

func (o *OperationAdd) Exe() int {
	//return o.oper.num1 + o.oper.num2
	return 1
}

func (o *OperationSub) Exe() int {
	//return o.oper.num1 - o.oper.num2
	return 2
}

func (o *OperationMul) Exe() int {
	//return o.oper.num1 * o.oper.num2
	return 3
}

func (o *OperationDiv) Exe() int {
	//return o.oper.num1 / o.oper.num2
	return 4
}
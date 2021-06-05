package calc

type OperationInterface interface {
	Exe(a, b int) int
}

type OperationAdd struct {}
type OperationSub struct {}
type OperationMul struct {}
type OperationDiv struct {}

func (o *OperationAdd) Exe(a, b int) int {
	return a + b
}

func (o *OperationSub) Exe(a, b int) int {
	return a - b
}

func (o *OperationMul) Exe(a, b int) int {
	return a * b
}

func (o *OperationDiv) Exe(a, b int) int {
	return a / b
}
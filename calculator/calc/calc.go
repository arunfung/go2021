package calc

type OperationInterface interface {
	Exe(a, b int) int
}

type OAdd struct {}
type OSub struct {}
type OMul struct {}
type ODiv struct {}

func (o *OAdd) Exe(a, b int) int {
	return a + b
}

func (o *OSub) Exe(a, b int) int {
	return a - b
}

func (o *OMul) Exe(a, b int) int {
	return a * b
}

func (o *ODiv) Exe(a, b int) int {
	return a / b
}
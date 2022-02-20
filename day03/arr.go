package main

type Arr struct {
	name string
}

func (a *Arr) GetName() string {
	return a.name
}

type MySlice struct {
	ptr *[10]int
	len int
	cap int
	arr Arr
}

func NewMySlice(len, cap int) *MySlice {
	var ar [10]int
	return &MySlice{
		ptr: &ar,
		len: len,
		cap: cap,
		arr: Arr{
			name: "arunfung",
		},
	}
}

func (m *MySlice) GetLen() int {
	return m.len
}

type name interface {
}

package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	name string
	age  int
}

func main() {
	t := Stu{"skidoo", 10}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f)
	}

	//var v int = 10
	//fmt.Println(v)
	//ref(v)
}

func ref(i interface{}) {
	rt := reflect.TypeOf(i)
	rv := reflect.ValueOf(i)
	fmt.Printf("rt type is %s \n", rt)
	fmt.Printf("rt type is %T \n", rv)

	rvi := rv.Interface()
	rtv := reflect.ValueOf(rt)

	fmt.Printf("rtv type is %T \n", rtv.Interface())
	fmt.Printf("rvi type is %T \n", rvi)
}

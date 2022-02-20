package edu_test

import (
	"encoding/json"
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

type Config struct {
	DataPath string `json:"data_path"`
	BasePath string `json:"base_path"`
}

func TestJson(t *testing.T) {
	c := Config{"DataPath", "BasePath"}
	cj, _ := json.Marshal(c)
	fmt.Println(string(cj))
	cj1, _ := json.MarshalIndent(c, "", "    ")
	fmt.Println(string(cj1))

	js := `{"data_path":"路径","base_path":"路径"}`
	cb := []byte(js)
	var c1 Config
	json.Unmarshal(cb, &c1)

	fmt.Println(c1)
}

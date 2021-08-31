package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ReadfileFun func(a interface{})

func ReadJosn(filepath string, a interface{}) error {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return err
	}
	err = json.Unmarshal(bytes, a)
	if err != nil {
		fmt.Println("解析数据失败", err)
		return err
	}
	return nil
}
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, nil
}

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	file := "file/1.txt"
	data := "arun"
	go Oper(file, data)
	go Oper(file, data)
	go Oper(file, data)
	go Oper(file, data)
	time.Sleep(1e9)
}
func Oper(path, data string) {
	d, _ := ReadFile(path)
	if d == "" {
		WriteFile(path, data)
		fmt.Println("写信息成功 : ", data)
	} else {
		fmt.Println("存在信息", d)
	}
}

func ReadFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return "", err
	}
	return string(bytes), nil
}

func WriteFile(path, data string) (bool, error) {
	// 打开文件
	// 0666 是文件的写入权限
	outputFile, outputError := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if outputError != nil {
		return false, outputError
	}
	defer outputFile.Close()
	// 创建写的缓冲区
	outputWriter := bufio.NewWriter(outputFile)
	// 写入信息
	outputWriter.WriteString(data)
	// 刷新
	outputWriter.Flush()
	return true, nil
}

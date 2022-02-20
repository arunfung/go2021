package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputReader *bufio.Reader

type Cfunc func() (bool, error)

func init() {
	inputReader = bufio.NewReader(os.Stdin)
}

func (c Cfunc) call() (bool, error) {
	return c()
}

func CRead() string {
	input, _ := inputReader.ReadString('\n')
	return strings.TrimSpace(strings.TrimSuffix(input, "\n"))
}

// 格式化输出
func CReturn(a Cfunc) bool {
	fmt.Println("============== start ============== >>>>>>>>>>")
	// flag, err := a.call()
	flag, err := a()
	if err != nil {
		fmt.Println("错误信息：", err)
	}
	fmt.Println("============== end   ============== >>>>>>>>>>")
	fmt.Println("")
	return flag
}

// 输出指令
func Coper(operate []string) {
	// operate := [3]string{0: "登入", 1: "注册", 2: "退出"}
	for key, value := range operate {
		fmt.Printf("(%d) : %s \n", key, value)
	}
	fmt.Println("退出请输 x ")
}

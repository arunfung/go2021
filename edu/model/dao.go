package model

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Model interface {
	ToString() string
}

var (
	path   string                 = "/Users/arun/go/src/go2021/edu/data/" // 数据路径
	suffix string                 = ".sql"
	models map[string]interface{} // 记录标识 user =》 结构体
)

func init() {
	// 标识绑定注册
	models = make(map[string]interface{})
	models["user"] = &User{}
}

func rfdata(name, pirmay string, datas map[string]Model) error {
	f, fErr := os.Open(path + name + suffix)
	if fErr != nil {
		fmt.Println("文件读取异常,", fErr)
		return errors.New("文件查询不到 error")
	}
	// 关闭文件的资源流
	defer f.Close()
	// 创建读取的文件的缓冲区
	buf := bufio.NewReader(f)
	// 2. 遍历数据  每一行的数据 字段根据 , 分割；数据通过 \n 分割
	field := make([]string, 0)
	for {
		row, rerr := buf.ReadBytes('\n')
		if rerr != nil {
			if rerr == io.EOF { // 是否文件读取结束
				break
			}
			fmt.Println("抛出缓存读取文件异常", rerr)
		}
		// 去掉字符串，并分割数据
		data := strings.Split(strings.TrimSuffix(string(row), "\n"), ",")
		//fmt.Println("读取到的数据", data)

		// 根据数据判断操作 ； 是记录字段还是设置数据
		if len(field) == 0 {
			field = data // 存储字段
		} else {}
	}
	fmt.Println("存储的数据", field)
	return nil
}
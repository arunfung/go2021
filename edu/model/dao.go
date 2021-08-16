package model

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Model interface {
	ToString() string
}

var (
	path   string                 = "/Users/arun/go/src/arunfung/go2021/edu/data/" // 数据路径
	suffix string                 = ".sql"
	models map[string]interface{} // 记录标识 user =》 结构体
)

func init() {
	// 标识绑定注册
	models = make(map[string]interface{})
	models["user"] = NewUser

	userDatas = make(map[string]Model, 0)
	rfdata("user", "username", userDatas)
}

func rfdata(name, primary string, datas map[string]Model) error {
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
			for k, v := range data {
				field[k] = strings.TrimSpace(strings.TrimSuffix(v, "\n"))
			}
		} else {
			toModel(name, primary, datas, data, field)
		}
	}
	//fmt.Println("存储的数据", field)
	return nil
}

func toModel(name string, primary string, datas map[string]Model, data []string, field []string) error {
	if models[name] == nil {
		return errors.New("不存在模型 : " + name)
	}
	//
	modelV := reflect.ValueOf(models[name]).Call([]reflect.Value{})[0]
	//fmt.Printf("modelV type %T \n", modelV)
	//fmt.Println("data 数据", data)
	//fmt.Println("field 数据", field)

	var primaryValue string
	for k, v := range data {
		if field[k] == primary { // 判断是否为主键字段的值 -》
			primaryValue = v
			//fmt.Println("查询的主键值 : ", primaryValue)
		}
		fset := modelV.MethodByName("Set" + strings.Title(field[k]))
		//fmt.Printf("fset type %T \n", fset)
		//fmt.Println("fset ", fset)

		fset.Call([]reflect.Value{
			reflect.ValueOf(toTypeValue(modelV, field[k], v)),
		})
	}
	//fmt.Println("model", modelV)
	datas[primaryValue] = modelV.Interface().(Model)

	return nil
}

func toTypeValue(modelV reflect.Value, field string, value string) interface{} {
	mtype := reflect.Zero(modelV.Type().Elem()).FieldByName(field).Type().Name()
	switch mtype {
	case "int":
		b, _ := strconv.Atoi(value)
		return b
	}
	return string(value)
}
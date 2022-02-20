package calc

import (
	"fmt"
	"reflect"
)

var opers map[string]interface{}

func init() {
	opers = make(map[string]interface{}, 0)
	opers["+"] = NewAdd
	opers["-"] = NewSub
	opers["*"] = NewMul
	opers["/"] = NewDiv
}

func OperationFactory(num1, num2 int, flag string) OperationInter {
	oper := opers[flag]
	// 根据类型得到反射的结构体对象Value
	valueOper := reflect.ValueOf(oper) // reflect.value
	fmt.Println("reflect.ValueOf(oper)", valueOper)
	fmt.Printf("type:%T \n", valueOper)

	// 设置调用函数的参数
	args := []reflect.Value{
		reflect.ValueOf(num1), reflect.ValueOf(num2),
	}

	// 根据类型调用函数或方法
	arrValueOper := valueOper.Call(args)[0]

	fmt.Println("valueOper.Call", arrValueOper)
	// reflect.Value => 解析为原来的对象 OperationInterface
	fmt.Printf("type:%T \n", arrValueOper)

	// 利用接口interface{} 断言转变为原有类型 ； interface{}=> 具体类型
	opertionin := arrValueOper.Interface().(OperationInter) // 返回一个 interface{} 类型

	fmt.Printf("retopertionin-type:%T \n", opertionin)
	return opertionin

}

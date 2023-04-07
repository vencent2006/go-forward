package reflect

import (
	"reflect"
)

func IterateFunc(entity any) (map[string]FuncInfo, error) {
	typ := reflect.TypeOf(entity)
	numMethod := typ.NumMethod()
	res := make(map[string]FuncInfo, numMethod)
	for i := 0; i < numMethod; i++ {
		method := typ.Method(i)
		fn := method.Func
		// input
		numIn := fn.Type().NumIn()
		input := make([]reflect.Type, 0, numIn)
		inputValues := make([]reflect.Value, 0, numIn)
		input = append(input, reflect.TypeOf(entity))              // 先把结构体类型传进来
		inputValues = append(inputValues, reflect.ValueOf(entity)) // 先把结构体类型的值传进来
		for i := 1; i < numIn; i++ {
			fnInType := fn.Type().In(i)
			input = append(input, fnInType)
			inputValues = append(inputValues, reflect.Zero(fnInType)) // 用零值发起调用
		}
		// output
		numOut := fn.Type().NumOut()
		output := make([]reflect.Type, 0, numOut)
		for i := 0; i < numOut; i++ {
			output = append(output, fn.Type().Out(i))
		}
		// result
		resValues := fn.Call(inputValues)
		result := make([]any, 0, len(resValues))
		for _, v := range resValues {
			result = append(result, v.Interface()) // v.Interface很重要
		}
		res[method.Name] = FuncInfo{
			Name:        method.Name,
			InputTypes:  input,
			OutputTypes: output,
			Result:      result,
		}
	}
	return res, nil
}

type FuncInfo struct {
	Name        string         // 函数名
	InputTypes  []reflect.Type // 参数
	OutputTypes []reflect.Type // 返回值
	Result      []any          // 调用的结果
}

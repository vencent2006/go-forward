package main

import (
	"fmt"
	"reflect"
)

func MyAdd(a int, b int) int {
	return a + b
}

func MyAdd2(a int, b int) int {
	return a - b
}

func CallAdd(f func(a int, b int) int) {
	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		// 不是方法
		return
	}

	// 构造argv
	argv := make([]reflect.Value, 2)
	argv[0] = reflect.ValueOf(1)
	argv[1] = reflect.ValueOf(2)

	// 只要反射包的类型是reflect.Func类型
	result := v.Call(argv)
	fmt.Println(result[0].Int())
}

func main() {
	CallAdd(MyAdd)  // 3
	CallAdd(MyAdd2) // -1
}

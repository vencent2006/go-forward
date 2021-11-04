/**
 * @Author: vincent
 * @Description:
 * @File:  reflection_test
 * @Version: 1.0.0
 * @Date: 2021/11/4 09:12
 */

package reflection_demo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValueOf(t *testing.T) {
	// 声明整型变量a并赋初值
	var a int = 1024

	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)

	// 获取interface{}类型的值, 通过类型断言转换
	var getA int = valueOfA.Interface().(int)

	// 获取64位的值, 强制类型转换为int类型
	var getA2 int = int(valueOfA.Int())

	fmt.Println(getA, getA2)
}

func TestValueOf2(t *testing.T) {
	type Turbo struct {
		Name string
		Age  int
	}
	turbo := &Turbo{
		Name: "迈莫coding",
		Age:  1,
	}
	value := reflect.ValueOf(turbo)
	fmt.Println("value is ", value)
	fmt.Println("value.Kind() is ", value.Kind())
	if value.Kind() == reflect.Ptr {
		fmt.Println("value.Elem is ", value.Elem())
		value = value.Elem()
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fmt.Printf("字段类型：%v, 字段值：%v\n", field.Type(), field.Interface())
		}

		fmt.Println("--- 获取Name字段的field")
		st := value.FieldByName("Name")
		fmt.Printf("Name field is %v\n", st.Interface())
	}
}

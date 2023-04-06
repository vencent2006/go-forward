package reflect

import (
	"errors"
	"reflect"
)

// IterateFields 遍历字段
func IterateFields(entity any) (map[string]any, error) {
	if entity == nil {
		return nil, errors.New("不支持 nil")
	}
	typ := reflect.TypeOf(entity)
	val := reflect.ValueOf(entity)
	if val.IsZero() {
		return nil, errors.New("不支持零值") // 零值：对象占据的内存空间比特位都为0
	}

	for typ.Kind() == reflect.Pointer {
		// 拿到指针指向的对象
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil, errors.New("不支持类型")
	}

	numField := typ.NumField()
	res := make(map[string]any, numField)
	for i := 0; i < numField; i++ {
		fieldType := typ.Field(i)  // 字段的类型
		fieldValue := val.Field(i) // 字段的值
		if fieldType.IsExported() {
			res[fieldType.Name] = fieldValue.Interface()
		} else {
			res[fieldType.Name] = reflect.Zero(fieldType.Type).Interface()
		}
	}
	return res, nil
}

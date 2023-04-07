package reflect

import (
	"errors"
	"reflect"
)

func IterateArrayOrSlice(entity any) ([]any, error) {
	val := reflect.ValueOf(entity)
	if val.Type().Kind() != reflect.Array && val.Type().Kind() != reflect.Slice {
		return nil, errors.New("类型不是Array和Slice")
	}
	length := val.Len()
	res := make([]any, 0, length)
	for i := 0; i < length; i++ {
		res = append(res, val.Index(i).Interface())
	}
	return res, nil
}

// keys, values, error
func IterateMap(entity any) ([]any, []any, error) {
	val := reflect.ValueOf(entity)
	if val.Type().Kind() != reflect.Map {
		return nil, nil, errors.New("类型不是Map")
	}
	resKeys := make([]any, 0, val.Len())
	resValues := make([]any, 0, val.Len())
	keys := val.MapKeys()
	for _, key := range keys {
		v := val.MapIndex(key)
		resKeys = append(resKeys, key.Interface())
		resValues = append(resValues, v.Interface())
	}
	return resKeys, resValues, nil
}

// 采用MapRange返回的iterator来遍历
func IterateMap2(entity any) ([]any, []any, error) {
	val := reflect.ValueOf(entity)
	if val.Type().Kind() != reflect.Map {
		return nil, nil, errors.New("类型不是Map")
	}
	resKeys := make([]any, 0, val.Len())
	resValues := make([]any, 0, val.Len())
	it := val.MapRange()
	for it.Next() {
		resKeys = append(resKeys, it.Key().Interface())
		resValues = append(resValues, it.Value().Interface())
	}
	return resKeys, resValues, nil
}

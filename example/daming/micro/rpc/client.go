package rpc

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

func InitClientProxy(service Service) error {
	// 在这里初始化一个proxy
	return setFuncField(service, nil)
}

func setFuncField(service Service, p Proxy) error {
	if service == nil {
		return errors.New("rpc: 不支持nil")
	}

	val := reflect.ValueOf(service)
	typ := val.Type()
	if typ.Kind() != reflect.Pointer || val.Elem().Kind() != reflect.Struct {
		return errors.New("rpc: 只支持指向结构体的一级指针")
	}

	val = val.Elem()
	typ = typ.Elem()

	numField := typ.NumField()
	for i := 0; i < numField; i++ {
		fieldTyp := typ.Field(i)
		fieldVal := val.Field(i)
		// 我要设置值给 GetById
		if fieldVal.CanSet() {
			// 这个地方才是真正的将本地调用捕捉到的地方
			fn := func(args []reflect.Value) (results []reflect.Value) {
				// args[0] 是context; context是不会传到服务端的，但是context的内容，可能会传到服务端
				ctx := args[0].Interface().(context.Context)
				// args[1] 是req
				req := &Request{
					ServiceName: service.Name(),
					MethodName:  fieldTyp.Name,
					Arg:         args[1].Interface(),
				}

				// reflect.New返回的是一个指针
				retVal := reflect.New(fieldTyp.Type.Out(0)).Elem()
				//fmt.Println(req)
				// 要真的发起调用
				resp, err := p.Invoke(ctx, req)
				if err != nil {
					return []reflect.Value{retVal, reflect.ValueOf(err)}
				}

				// 这里怎么办
				fmt.Println(resp)
				return []reflect.Value{retVal, reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}
			fnVal := reflect.MakeFunc(fieldTyp.Type, fn)
			fieldVal.Set(fnVal)
		}
	}
	return nil
}

// 生成 proxy的mock
// mockgen -destination=micro/rpc/mock_proxy_gen_test.go -package=rpc -source=micro/rpc/types.go Proxy

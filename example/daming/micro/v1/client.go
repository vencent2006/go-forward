package v1

import (
	"context"
	"encoding/json"
	"errors"
	"example/daming/micro/v1/message"
	"github.com/silenceper/pool"
	"net"
	"reflect"
	"time"
)

func InitClientProxy(addr string, service Service) error {
	client, err := NewClient(addr)
	if err != nil {
		return err
	}
	// 在这里初始化一个proxy
	return setFuncField(service, client)
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
				retVal := reflect.New(fieldTyp.Type.Out(0).Elem()) // reflect.New返回的是一个指针

				// args[0] 是context; context是不会传到服务端的，但是context的内容，可能会传到服务端
				ctx := args[0].Interface().(context.Context)
				// args[1] 是req
				reqData, err := json.Marshal(args[1].Interface())
				if err != nil {
					return []reflect.Value{retVal, reflect.ValueOf(err)}
				}
				req := &message.Request{
					ServiceName: service.Name(),
					MethodName:  fieldTyp.Name,
					Data:        reqData,
				}

				// 要真的发起调用
				resp, err := p.Invoke(ctx, req)
				if err != nil {
					return []reflect.Value{retVal, reflect.ValueOf(err)}
				}

				err = json.Unmarshal(resp.Data, retVal.Interface())
				if err != nil {
					return []reflect.Value{retVal, reflect.ValueOf(err)}
				}

				// 这里怎么办
				//fmt.Println(resp)
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

type Client struct {
	pool pool.Pool
}

func NewClient(addr string) (*Client, error) {
	p, err := pool.NewChannelPool(&pool.Config{
		InitialCap: 1,
		MaxCap:     30,
		MaxIdle:    10,
		Factory: func() (interface{}, error) {
			return net.DialTimeout("tcp", addr, time.Second*3)
		},
		Close: func(i interface{}) error {
			return i.(net.Conn).Close()
		},
		Ping:        nil,
		IdleTimeout: time.Minute * 1,
	})
	if err != nil {
		return nil, err
	}
	return &Client{pool: p}, nil
}

func (c *Client) Invoke(ctx context.Context, req *message.Request) (*message.Response, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 正儿八经地把请求发过去服务端
	resp, err := c.Send(data)
	if err != nil {
		return nil, err
	}
	return &message.Response{Data: resp}, nil
}

func (c *Client) Send(data []byte) ([]byte, error) {
	// todo pool.Put() 在什么地方调用呢？
	val, err := c.pool.Get()
	if err != nil {
		return nil, err
	}
	conn := val.(net.Conn)
	defer func() {
		_ = conn.Close()
	}()

	res := EncodeMsg(data)
	_, err = conn.Write(res)
	if err != nil {
		return nil, err
	}

	return ReadMsg(conn)
}

package rpc

import (
	"context"
	"errors"
	"example/daming/micro/rpc/message"
	"example/daming/micro/rpc/serialize"
	"example/daming/micro/rpc/serialize/json"
	"github.com/silenceper/pool"
	"net"
	"reflect"
	"strconv"
	"time"
)

// 要为 GetById之类的函数类型的字段赋值
func (c *Client) InitService(service Service) error {
	// 在这里初始化一个proxy
	return setFuncField(service, c, c.serializer)
}

func setFuncField(service Service, p Proxy, s serialize.Serializer) error {
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
				reqData, err := s.Encode(args[1].Interface()) // 序列化
				if err != nil {
					return []reflect.Value{retVal, reflect.ValueOf(err)}
				}

				meta := make(map[string]string, 2)
				// 我确实设置了超时
				if deadline, ok := ctx.Deadline(); ok {
					meta[META_KEY_DEADLINE] = strconv.FormatInt(deadline.UnixMilli(), 10)
				}
				// one way
				if isOneway(ctx) {
					meta[META_KEY_ONEWAY] = META_VAL_ONEWAY_TRUE
				}
				req := &message.Request{
					ServiceName: service.Name(),
					MethodName:  fieldTyp.Name,
					Data:        reqData,
					Serializer:  s.Code(), // 序列化类型
					Meta:        meta,
				}
				req.CalculateHeaderLength()
				req.CalculateBodyLength()

				// 要真的发起调用
				resp, err := p.Invoke(ctx, req)
				if err != nil {
					// 反序列化的error
					return []reflect.Value{retVal, reflect.ValueOf(err)}
				}

				var retErr error
				if len(resp.Error) > 0 {
					// 服务端传过来的 error
					retErr = errors.New(string(resp.Error))
				}

				if len(resp.Data) > 0 {
					err = s.Decode(resp.Data, retVal.Interface()) // 反序列化
					if err != nil {
						// 反序列化的 error
						return []reflect.Value{retVal, reflect.ValueOf(err)}
					}
				}

				var retErrVal reflect.Value
				if retErr == nil {
					retErrVal = reflect.Zero(reflect.TypeOf(new(error)).Elem())
				} else {
					retErrVal = reflect.ValueOf(retErr)
				}

				// 这里怎么办
				//fmt.Println(resp)
				return []reflect.Value{retVal, retErrVal}
			}
			// 我要设置值给 GetById
			fnVal := reflect.MakeFunc(fieldTyp.Type, fn)
			fieldVal.Set(fnVal)
		}
	}
	return nil
}

// 生成 proxy的mock
// mockgen -destination=micro/rpc/mock_proxy_gen_test.go -package=rpc -source=micro/rpc/types.go Proxy

type Client struct {
	pool       pool.Pool
	serializer serialize.Serializer
}

type ClientOption func(client *Client)

func ClientWithSerializer(sl serialize.Serializer) ClientOption {
	return func(client *Client) {
		client.serializer = sl
	}
}

func NewClient(addr string, opts ...ClientOption) (*Client, error) {
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
	res := &Client{
		pool:       p,
		serializer: &json.Serializer{},
	}
	for _, opt := range opts {
		opt(res)
	}
	return res, nil
}

func (c *Client) Invoke(ctx context.Context, req *message.Request) (*message.Response, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	ch := make(chan struct{})
	var (
		resp *message.Response
		err  error
	)
	go func() {
		defer close(ch)
		resp, err = c.doInvoke(ctx, req)
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ch:
		return resp, err
	}
}

func (c *Client) doInvoke(ctx context.Context, req *message.Request) (*message.Response, error) {
	data := message.EncodeReq(req)
	// 正儿八经地把请求发过去服务端
	resp, err := c.send(ctx, data)
	if err != nil {
		return nil, err
	}
	return message.DecodeResp(resp), nil
}

func (c *Client) send(ctx context.Context, data []byte) ([]byte, error) {
	// todo pool.Put() 在什么地方调用呢？
	val, err := c.pool.Get()
	if err != nil {
		return nil, err
	}
	conn := val.(net.Conn)
	defer func() {
		_ = conn.Close()
	}()

	//res := EncodeMsg(data)
	_, err = conn.Write(data) // 已经encode过了
	if err != nil {
		return nil, err
	}

	if isOneway(ctx) {
		// 当然 err可以等于nil，
		return nil, errors.New("micro: 这是一个 oneway 调用， 你不应该处理任何结果")
	}

	return ReadMsg(conn)
}

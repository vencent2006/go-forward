package rpc

import (
	"context"
	"errors"
	"example/daming/micro/rpc/message"
	"example/daming/micro/rpc/serialize"
	"example/daming/micro/rpc/serialize/json"
	"net"
	"reflect"
	"strconv"
	"time"
)

// Server可以是一个Proxy，在服务端的proxy
type Server struct {
	services    map[string]reflectionStub
	serializers map[uint8]serialize.Serializer // 服务端可能会有多个序列化协议
}

func NewServer() *Server {
	res := &Server{
		services:    make(map[string]reflectionStub, 16),     // 先预估个容量，比如16
		serializers: make(map[uint8]serialize.Serializer, 4), // 先预估个容量，比如4
	}
	res.RegisterSerializer(&json.Serializer{})
	return res
}

func (s *Server) RegisterSerializer(serializer serialize.Serializer) {
	s.serializers[serializer.Code()] = serializer
}

func (s *Server) RegisterService(service Service) {
	s.services[service.Name()] = reflectionStub{
		s:           service,
		value:       reflect.ValueOf(service),
		serializers: s.serializers, // 先预估个容量，比如4
	}
}

func (s *Server) Start(network, addr string) error {
	listener, err := net.Listen(network, addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			// 一般是端口被占用的错误
			return err
		}

		go func() {
			if err := s.handleConn(conn); err != nil {
				_ = conn.Close()
			}
		}()
	}
}

// 我们可以认为一个请求包含2个部分
// 1. 长度字段：8个字节
// 2. 请求数据
// 响应也是这个规范
func (s *Server) handleConn(conn net.Conn) error {
	for {
		reqBs, err := ReadMsg(conn)
		if err != nil {
			return err
		}

		// 还原调用信息
		req := message.DecodeReq(reqBs)
		ctx := context.Background()
		cancel := func() {}
		if deadlineStr, ok := req.Meta[META_KEY_DEADLINE]; ok {
			// todo 如果er != nil， 那么是要返回默认超时时间还是直接返回客户端错误呢
			if deadline, er := strconv.ParseInt(deadlineStr, 10, 64); er == nil {
				ctx, cancel = context.WithDeadline(ctx, time.UnixMilli(deadline))
			}
		}
		oneway, ok := req.Meta[META_KEY_ONEWAY]
		if ok && oneway == META_VAL_ONEWAY_TRUE {
			ctx = CtxWithOneWay(ctx)
		}
		resp, err := s.Invoke(ctx, req)
		cancel() // 用完了(ctx)，就马上cancel掉
		if err != nil {
			// 这个可能是你的业务 error
			resp.Error = []byte(err.Error())
		}
		resp.CalculateHeaderLength()
		resp.CalculateBodyLength()
		//res := EncodeMsg(resp.Data)
		_, err = conn.Write(message.EncodeResp(resp))
		if err != nil {
			return err
		}

	}
}

func (s *Server) Invoke(ctx context.Context, req *message.Request) (*message.Response, error) {
	// 还原了调用信息， 你已经知道：service name, method name 和 args了
	// 发起业务调用了
	service, ok := s.services[req.ServiceName]
	resp := &message.Response{
		RequestID:  req.RequestID,
		Version:    req.Version,
		Compresser: req.Compresser,
		Serializer: req.Serializer,
	}

	if !ok { // 不存在
		resp.Error = []byte("你要调用的服务不存在")
		return resp, nil
	}

	if isOneway(ctx) {
		go func() {
			_, _ = service.invoke(ctx, req)
		}()
		// 为防止中间件以为这是个error，也可以直接返回nil
		return nil, errors.New("micro: 微服务服务端 oneway 请求")
	}

	respData, err := service.invoke(ctx, req)
	if err != nil {
		return resp, err
	}
	resp.Data = respData

	return resp, nil
}

// 使用 reflectionStub做了一个和client对称的proxy
// 即：server和client都是通过proxy来实现rpc的调用代理
// 当前的代理是reflection(反射)的代理，如果未来要用unsafe的代理，就直接改stub就行了
type reflectionStub struct {
	s           Service
	value       reflect.Value
	serializers map[uint8]serialize.Serializer // 服务端可能会有多个序列化协议
}

func (s *reflectionStub) invoke(ctx context.Context, req *message.Request) ([]byte, error) {
	// 反射找到方法，并且执行调用
	method := s.value.MethodByName(req.MethodName)
	in := make([]reflect.Value, 2) // 一共2个参数，一个是context，一个是request
	in[0] = reflect.ValueOf(ctx)
	inReq := reflect.New(method.Type().In(1).Elem()) // reflect.New返回的是一个指针; 注意这里Elem()的位置，必须得放入到reflect.New里才能分配内存
	serializer, ok := s.serializers[req.Serializer]
	if !ok {
		return nil, errors.New("micro: 不支持的序列化协议")
	}
	err := serializer.Decode(req.Data, inReq.Interface()) // 反序列化
	if err != nil {
		return nil, err
	}
	in[1] = inReq
	results := method.Call(in)
	// result[0] 是返回值
	// result[1] 是error
	if results[1].Interface() != nil {
		return nil, results[1].Interface().(error)
	}

	return serializer.Encode(results[0].Interface()) // 序列化
}

//
//func (s *Server) handleMsg(reqData []byte) ([]byte, error) {
//
//	// 还原了调用信息， 你已经知道：service name, method name 和 args了
//	// 发起业务调用了
//	service, ok := s.services[req.ServiceName]
//	if !ok { // 不存在
//		return nil, errors.New("你要调用的服务不存在")
//	}
//	// 反射找到方法，并且执行调用
//	val := reflect.ValueOf(service)
//	method := val.MethodByName(req.MethodName)
//	in := make([]reflect.Value, 2) // 一共2个参数，一个是context，一个是request
//	// 暂时我们不知道怎么传这个 context，所以我们就直接写死
//	in[0] = reflect.ValueOf(context.Background())
//	inReq := reflect.New(method.Type().In(1).Elem()) // reflect.New返回的是一个指针; 注意这里Elem()的位置，必须得放入到reflect.New里才能分配内存
//	err = json.Unmarshal(req.Data, inReq.Interface())
//	if err != nil {
//		return nil, err
//	}
//	in[1] = inReq
//	results := method.Call(in)
//	// result[0] 是返回值
//	// result[1] 是error
//	if results[1].Interface() != nil {
//		return nil, results[1].Interface().(error)
//	}
//	resp, err := json.Marshal(results[0].Interface())
//	return resp, err
//}

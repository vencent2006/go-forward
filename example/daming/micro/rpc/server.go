package rpc

import (
	"context"
	"encoding/json"
	"example/daming/micro/rpc/message"
	"net"
	"reflect"
)

// Server可以是一个Proxy，在服务端的proxy
type Server struct {
	services map[string]reflectionStub
}

func NewServer() *Server {
	return &Server{
		services: make(map[string]reflectionStub, 16), // 先预估个容量，比如16
	}
}

func (s *Server) RegisterService(service Service) {
	s.services[service.Name()] = reflectionStub{
		s:     service,
		value: reflect.ValueOf(service),
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
		resp, err := s.Invoke(context.Background(), req)
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

	respData, err := service.invoke(ctx, req.MethodName, req.Data)
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
	s     Service
	value reflect.Value
}

func (s *reflectionStub) invoke(ctx context.Context, methodName string, data []byte) ([]byte, error) {
	// 反射找到方法，并且执行调用
	method := s.value.MethodByName(methodName)
	in := make([]reflect.Value, 2) // 一共2个参数，一个是context，一个是request
	// 暂时我们不知道怎么传这个 context，所以我们就直接写死
	in[0] = reflect.ValueOf(context.Background())
	inReq := reflect.New(method.Type().In(1).Elem()) // reflect.New返回的是一个指针; 注意这里Elem()的位置，必须得放入到reflect.New里才能分配内存
	err := json.Unmarshal(data, inReq.Interface())
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

	return json.Marshal(results[0].Interface())
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

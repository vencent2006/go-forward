package rpc

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"reflect"
)

type Server struct {
	services map[string]Service
}

func NewServer() *Server {
	return &Server{
		services: make(map[string]Service, 16), // 先预估个容量，比如16
	}
}

func (s *Server) RegisterService(service Service) {
	s.services[service.Name()] = service
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
		// lenBs 是长度字段的字节表示
		lenBs := make([]byte, numOfLengthBytes)
		_, err := conn.Read(lenBs)
		if err != nil {
			return err
		}

		// 消息有多长？
		length := binary.BigEndian.Uint64(lenBs)
		reqBs := make([]byte, length)
		n, err := conn.Read(reqBs)
		if err != nil {
			return err
		}
		if uint64(n) != length {
			return fmt.Errorf("micro: 数据读取不完整")
		}
		respData, err := s.handleMsg(reqBs)
		if err != nil {
			// 这个可能是你的业务 error
			// todo 暂时不知道怎么回传error， 所以我们简单记录一下
			return err
		}
		respLen := len(respData)
		// 我要在这，构建响应数据
		// data = respLen 的64位表示 + respData
		res := make([]byte, respLen+numOfLengthBytes)

		// 第一步: 把长度写进去前 numOfLengthBytes 个字节
		binary.BigEndian.PutUint64(res[:numOfLengthBytes], uint64(respLen))
		// 第二步: 写入数据
		copy(res[numOfLengthBytes:], respData)
		_, err = conn.Write(res)
		if err != nil {
			return err
		}

	}
}

func (s *Server) handleMsg(reqData []byte) ([]byte, error) {
	// 还原调用信息
	req := &Request{}
	err := json.Unmarshal(reqData, req)
	if err != nil {
		return nil, err
	}

	// 还原了调用信息， 你已经知道：service name, method name 和 args了
	// 发起业务调用了
	service, ok := s.services[req.ServiceName]
	if !ok { // 不存在
		return nil, errors.New("你要调用的服务不存在")
	}
	// 反射找到方法，并且执行调用
	val := reflect.ValueOf(service)
	method := val.MethodByName(req.MethodName)
	in := make([]reflect.Value, 2) // 一共2个参数，一个是context，一个是request
	// 暂时我们不知道怎么传这个 context，所以我们就直接写死
	in[0] = reflect.ValueOf(context.Background())
	inReq := reflect.New(method.Type().In(1).Elem()) // reflect.New返回的是一个指针; 注意这里Elem()的位置，必须得放入到reflect.New里才能分配内存
	err = json.Unmarshal(req.Arg, inReq.Interface())
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
	resp, err := json.Marshal(results[0].Interface())
	return resp, err
}
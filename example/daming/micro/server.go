package micro

import (
	"context"
	"example/daming/micro/registry"
	"google.golang.org/grpc"
	"net"
	"time"
)

type ServerOption func(server *Server)

type Server struct {
	name            string
	registry        registry.Registry
	registryTimeout time.Duration
	*grpc.Server    // 使用组合类型
	listener        net.Listener
}

func NewServer(name string, opts ...ServerOption) (*Server, error) {
	res := &Server{
		name:            name,
		Server:          grpc.NewServer(),
		registryTimeout: time.Second * 10, // 默认10s超时
	}
	for _, opt := range opts {
		opt(res)
	}
	return res, nil
}

// Start 当用户调用这个方法的时候，就是服务已经准备好
func (s *Server) Start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.listener = listener

	// 有注册中心，要注册了
	if s.registry != nil {
		//  在这里注册
		ctx, cancel := context.WithTimeout(context.Background(), s.registryTimeout)
		defer cancel()
		err = s.registry.Register(ctx, registry.ServiceInstance{
			Name: s.name,
			// 你的定位信息从哪里来
			Address: listener.Addr().String(), // todo 这种写法在容器里会有问题

		})
		if err != nil {
			return err
		}
		// 这里已经注册成功了
		//defer func() {
		//	// 忽略 或者 log下错误
		//	_ = s.registry.Close()
		//}()
	}

	err = s.Serve(s.listener)

	return err
}

func (s *Server) Close() error {
	if s.registry != nil {
		err := s.registry.Close()
		if err != nil {
			return err
		}
	}
	s.GracefulStop() // 可以看下GracefulStop
	return nil
}

func ServerWithRegistry(r registry.Registry) ServerOption {
	return func(server *Server) {
		server.registry = r
	}
}

func ServerWithTimeout(t time.Duration) ServerOption {
	return func(server *Server) {
		server.registryTimeout = t
	}
}

package rpc

import "context"

const numOfLengthBytes = 8

type Service interface {
	Name() string
}

type Proxy interface {
	Invoke(ctx context.Context, req *Request) (*Response, error)
}

type Request struct {
	ServiceName string
	MethodName  string
	Arg         []byte
}

type Response struct {
	Data []byte
}

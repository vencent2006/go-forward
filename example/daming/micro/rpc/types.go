package rpc

import "context"

type Service interface {
	Name() string
}

type Proxy interface {
	Invoke(ctx context.Context, req *Request) (*Response, error)
}

type Request struct {
	HeadLength uint32
	BodyLength uint32
	RequestID  uint32
	Version    uint8
	Compresser uint8
	Serializer uint8

	ServiceName string
	MethodName  string

	Meta map[string]string

	Data []byte
}

type Response struct {
	HeadLength uint32
	BodyLength uint32
	RequestID  uint32
	Version    uint8
	Compresser uint8
	Serializer uint8

	Error []byte

	Data []byte
}

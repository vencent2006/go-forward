package v1

import (
	"context"
	"example/daming/micro/v1/message"
)

type Service interface {
	Name() string
}

type Proxy interface {
	Invoke(ctx context.Context, req *message.Request) (*message.Response, error)
}

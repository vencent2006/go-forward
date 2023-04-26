package rpc

import (
	"context"
	"example/daming/micro/rpc/message"
)

const (
	META_KEY_ONEWAY      = "one-way"
	META_KEY_DEADLINE    = "deadline"
	META_VAL_ONEWAY_TRUE = "true"
)

type Service interface {
	Name() string
}

type Proxy interface {
	Invoke(ctx context.Context, req *message.Request) (*message.Response, error)
}

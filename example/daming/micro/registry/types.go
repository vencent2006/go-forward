package registry

import (
	"context"
	"io"
)

type Registry interface {
	Register(ctx context.Context, si ServiceInstance) error
	UnRegister(ctx context.Context, si ServiceInstance) error
	ListServices(ctx context.Context, serviceName string) ([]ServiceInstance, error)
	Subscribe(serviceName string) (<-chan Event, error)
	// 显示的Close方法
	io.Closer
}

type ServiceInstance struct {
	Name string
	// Address 就是最关键的，定位信息
	Address string

	// 这边你可以任意加字段，完全取决于你的服务治理需要什么字段
}

type Event struct {
}

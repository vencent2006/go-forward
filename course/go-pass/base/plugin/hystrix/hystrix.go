

package hystrix

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/v3/client"
)

type clientWrapper struct {
	client.Client
}

//熔断逻辑
func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//正常执行
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(e error) error {
		//走熔断逻辑,每个服务都不一样
		fmt.Println(req.Service() + "." + req.Endpoint()+"的熔断逻辑")
		return e
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}

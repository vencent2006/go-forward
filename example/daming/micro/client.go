package micro

import (
	"context"
	"example/daming/micro/registry"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

type ClientOption func(c *Client)

type Client struct {
	insecure bool
	r        registry.Registry
	timeout  time.Duration
}

func NewClient(opts ...ClientOption) (*Client, error) {
	res := &Client{}
	for _, opt := range opts {
		opt(res)
	}
	return res, nil
}

func ClientInsecure() ClientOption {
	return func(c *Client) {
		c.insecure = true
	}
}

func ClientWithRegistry(r registry.Registry, timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.r = r
		c.timeout = timeout
	}
}

func (c *Client) Dial(ctx context.Context, service string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if c.r != nil {
		rb, err := NewRegistryBuilder(c.r, c.timeout)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithResolvers(rb))
	}
	if c.insecure {
		opts = append(opts, grpc.WithInsecure())
	}
	cc, err := grpc.DialContext(ctx, fmt.Sprintf("registry:///%s", service), opts...)
	return cc, err
}

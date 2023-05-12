package micro

import (
	"context"
	"example/daming/micro/registry"
	"google.golang.org/grpc/resolver"
	"time"
)

type grpcResolverBuilder struct {
	r       registry.Registry
	timeout time.Duration
}

func NewRegistryBuilder(r registry.Registry, timeout time.Duration) (*grpcResolverBuilder, error) {
	return &grpcResolverBuilder{r: r, timeout: timeout}, nil
}

func (b *grpcResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &grpcResolver{
		target:  target,
		cc:      cc,
		r:       b.r,
		close:   make(chan struct{}, 1),
		timeout: b.timeout,
	}

	r.resolve() // 要马上调用一次 resolve
	go r.watch()
	return r, nil
}

func (b *grpcResolverBuilder) Scheme() string {
	return "registry"
}

// grpcResolver
type grpcResolver struct {
	target  resolver.Target
	r       registry.Registry
	cc      resolver.ClientConn
	timeout time.Duration
	close   chan struct{}
}

func (g *grpcResolver) ResolveNow(options resolver.ResolveNowOptions) {
	g.resolve()
}

func (g *grpcResolver) watch() {
	events, err := g.r.Subscribe(g.target.Endpoint)
	if err != nil {
		g.cc.ReportError(err)
		return
	}
	for {
		select {
		case <-events:
			g.resolve() // 简单，幂等，但是给注册中心的压力要适当考虑;不然就要解析events的Type，但是会有时序一致性的问题
		case <-g.close:
			return
		}
	}
}

func (g *grpcResolver) resolve() {
	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()
	instances, err := g.r.ListServices(ctx, g.target.Endpoint)
	if err != nil {
		g.cc.ReportError(err)
		return
	}
	address := make([]resolver.Address, 0, len(instances))
	for _, si := range instances {
		address = append(address, resolver.Address{Addr: si.Address})
	}
	err = g.cc.UpdateState(resolver.State{Addresses: address})
	if err != nil {
		g.cc.ReportError(err)
		return
	}
}

func (g *grpcResolver) Close() {
	close(g.close)
}

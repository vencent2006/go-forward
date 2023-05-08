package grpc_resolver

import "google.golang.org/grpc/resolver"

type Builder struct {
}

func (b Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &Resolver{
		cc: cc,
	}
	// 要马上调用一次 ResolveNow
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}

func (b Builder) Scheme() string {
	return "registry"
}

type Resolver struct {
	cc resolver.ClientConn
}

func (r Resolver) ResolveNow(options resolver.ResolveNowOptions) {
	// 固定写死ip和端口
	// localhost:8081
	r.cc.UpdateState(resolver.State{
		Addresses: []resolver.Address{
			{
				Addr: "localhost:8081",
			},
		},
	})
}

func (r Resolver) Close() {
	// 没有资源要释放，什么也不干
}

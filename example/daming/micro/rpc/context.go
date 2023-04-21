package rpc

import "context"

type onewayKey struct{}

func CtxWithOneWay(ctx context.Context) context.Context {
	return context.WithValue(ctx, onewayKey{}, true)
}

func isOneway(ctx context.Context) bool {
	value := ctx.Value(onewayKey{})
	oneway, ok := value.(bool)
	return ok && oneway
}

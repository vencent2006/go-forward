package grpc

import (
	"context"
	"example/daming/micro/proto/gen"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	// grpc.WithInsecure, 表示不使用证书(https)
	cc, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	require.NoError(t, err)
	client := gen.NewUserServiceClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := client.GetById(ctx, &gen.GetByIdReq{Id: 13})
	require.NoError(t, err)
	t.Log(resp)

}

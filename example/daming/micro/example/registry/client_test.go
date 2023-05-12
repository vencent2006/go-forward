package registry

import (
	"context"
	"example/daming/micro"
	"example/daming/micro/proto/gen"
	"example/daming/micro/registry/etcd"
	"github.com/stretchr/testify/require"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	// grpc.WithInsecure, 表示不使用证书(https)
	// 格式是："scheme://[authority]/endpoint", scheme=registry, authority=空, endpoint=localhost:8081
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})
	require.NoError(t, err)
	r, err := etcd.NewRegistry(etcdClient)
	require.NoError(t, err)

	client, err := micro.NewClient(micro.ClientInsecure(), micro.ClientWithRegistry(r, time.Second*3))
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	cc, err := client.Dial(ctx, service_name)
	require.NoError(t, err)

	uc := gen.NewUserServiceClient(cc)
	resp, err := uc.GetById(ctx, &gen.GetByIdReq{Id: 13})
	require.NoError(t, err)
	t.Log(resp)
}

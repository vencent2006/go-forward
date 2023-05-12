package registry

import (
	"context"
	"example/daming/micro"
	"example/daming/micro/proto/gen"
	"example/daming/micro/registry/etcd"
	"fmt"
	"github.com/stretchr/testify/require"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})
	require.NoError(t, err)
	r, err := etcd.NewRegistry(etcdClient)
	require.NoError(t, err)
	us := &UserServiceServer{}
	server, err := micro.NewServer(service_name, micro.ServerWithRegistry(r), micro.ServerWithTimeout(time.Second*10))
	require.NoError(t, err)
	gen.RegisterUserServiceServer(server, us)
	// 我在这里调用 Start 方法，就意味着 us 已经完全准备好了
	err = server.Start(":8081")
	t.Log(err)
}

type UserServiceServer struct {
	gen.UnimplementedUserServiceServer
}

func (s UserServiceServer) GetById(ctx context.Context, req *gen.GetByIdReq) (*gen.GetByIdResp, error) {
	fmt.Println(req)
	return &gen.GetByIdResp{
		User: &gen.User{
			Name: "hello, world",
		},
	}, nil
}

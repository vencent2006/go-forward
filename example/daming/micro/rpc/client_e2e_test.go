package rpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestInitClientProxy(t *testing.T) {
	server := NewServer()
	server.RegisterService(&UserServiceServer{})
	serverAddr := ":8081"
	go func() {
		err := server.Start("tcp", serverAddr)
		t.Log(err)
	}()
	time.Sleep(time.Second * 3)
	usClient := &UserService{}
	err := InitClientProxy(serverAddr, usClient)
	require.NoError(t, err)
	resp, err := usClient.GetById(context.Background(), &GetByIdReq{Id: 123})
	require.NoError(t, err)
	assert.Equal(t, &GetByIdResp{Msg: "Hello, world"}, resp)
}

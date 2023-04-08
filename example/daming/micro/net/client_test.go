package net

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	go func() {
		err := Serve("tcp", ":8082")
		t.Log(err)
	}()

	time.Sleep(time.Second * 3)

	err := Connect("tcp", "localhost:8082")
	t.Log(err)

}

func TestClient_Send(t *testing.T) {
	go func() {
		err := NewServer().Start("tcp", ":8082")
		t.Log(err)
	}()

	time.Sleep(time.Second * 3)

	res, err := NewClient("tcp", "localhost:8082").Send("hello")
	require.NoError(t, err)
	assert.Equal(t, res, "hellohello")
}

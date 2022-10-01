/**
 * @Author: vincent
 * @Description:
 * @File:  echo
 * @Version: 1.0.0
 * @Date: 2022/9/30 16:01
 */

package tcp

import (
	"bufio"
	"context"
	"go-examples/course/go-redis-rewrite/go-redis/lib/logger"
	"go-examples/course/go-redis-rewrite/go-redis/lib/sync/atomic"
	"go-examples/course/go-redis-rewrite/go-redis/lib/sync/wait"
	"io"
	"net"
	"sync"
	"time"
)

type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

// Close EchoClient的客户端关闭
func (e *EchoClient) Close() error {
	// 等待10s的waitGroup
	e.Waiting.WaitWithTimeout(10 * time.Second)
	_ = e.Conn.Close()
	return nil
}

type EchoHandler struct {
	activeConn sync.Map       // 有多少client连接
	closing    atomic.Boolean // 是否正在关闭, 原子的boolean
}

func MakeHandler() *EchoHandler {
	return &EchoHandler{}
}

func (handler *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if handler.closing.Get() {
		// 准备关闭了，就不接收客户端连接了
		_ = conn.Close()
	}

	client := &EchoClient{
		Conn: conn,
	}

	handler.activeConn.Store(client, struct{}{}) // hashSet, value是空结构体
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// 结束符
				logger.Info("Connection close")
				handler.activeConn.Delete(client)
			} else {
				logger.Warn(err)
			}

			// 返回
			return
		}
		client.Waiting.Add(1) // 等会我，先别关闭我
		b := []byte(msg)
		_, _ = conn.Write(b)
		client.Waiting.Done()
	}
}

func (handler *EchoHandler) Close() error {
	logger.Info("handler shutting down")
	handler.closing.Set(true)
	// 关闭所有的client
	handler.activeConn.Range(func(key, value interface{}) bool {
		client := key.(*EchoClient)
		_ = client.Conn.Close()
		// 返回true，表示要遍历下一个
		return true
	})
	return nil
}

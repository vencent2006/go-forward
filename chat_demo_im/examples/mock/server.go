/**
 * @Author: vincent
 * @Description:
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2021/8/16 00:09
 */

package mock

import (
	"errors"
	kim "go-examples/chat_demo_im"
	"go-examples/chat_demo_im/logger"
	"go-examples/chat_demo_im/naming"
	"go-examples/chat_demo_im/tcp"
	"go-examples/chat_demo_im/websocket"
	"time"
)

type ServerDemo struct {
}

func (s *ServerDemo) Start(id, protocol, addr string) {
	var srv kim.Server

	service := &naming.DefaultServiceImpl{
		Id:       id,
		Protocol: protocol,
	}

	// 根据protocol来实例化
	if protocol == "ws" {
		srv = websocket.NewServer(addr, service)
	} else if protocol == "tcp" {
		srv = tcp.NewServer(addr, service)
	}

	// set handler
	handler := &ServerHandler{}
	srv.SetReadWait(time.Minute)
	srv.SetAcceptor(handler)
	srv.SetMessageListener(handler)
	srv.SetStateListener(handler)

	// start
	err := srv.Start()
	if err != nil {
		panic(err)
	}
}

type ServerHandler struct {
}

func (h *ServerHandler) Disconnect(channelId string) error {
	logger.Warnf("disconnect %s", channelId)
	return nil
}

func (h *ServerHandler) Receive(ag kim.Agent, payload []byte) {
	// 取payload
	ack := string(payload) + " from server "
	// 发送ack
	_ = ag.Push([]byte(ack))
}

func (h *ServerHandler) Accept(conn kim.Conn, wait time.Duration) (string, error) {
	// 1.读取: 客户端发送的鉴权数据包
	frame, err := conn.ReadFrame()
	if err != nil {
		return "", err
	}
	logger.Info("recv ", frame.GetOpCode())

	// 2.解析：数据包内容就是userId
	userId := string(frame.GetPayload())

	// 3.鉴权：这里只是为了示例，做一个fake验证，非空
	if userId == "" {
		return "", errors.New("user id is invalid(nil)")
	}

	return userId, nil

}

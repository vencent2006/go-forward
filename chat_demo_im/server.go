/**
 * @Author: vincent
 * @Description:
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2021/8/12 15:11
 */

package kim

import (
	"context"
	"go-examples/chat_demo_im/naming"
	"time"
)

type Server interface {
	// init
	naming.ServiceRegistration
	SetAcceptor(Acceptor)
	SetMessageListener(MessageListener)
	SetStateListener(StateListener)
	SetChannelMap(ChannelMap)
	SetReadWait(duration time.Duration)

	// runtime
	Start() error
	Shutdown(context context.Context) error

	// action
	Push(channelId string, data []byte) error
}

// interface group
type Acceptor interface {
	// Accept 返回一个握手完成的Channel对象或者一个error。
	// 业务层需要处理不同协议和网络环境的下连接握手协议
	Accept(conn Conn, wait time.Duration) (string, error)
}

type MessageListener interface {
	// 收到消息回调
	Receive(client Agent, buf []byte)
}

type StateListener interface {
	Disconnect(channelId string) error
}

/**
 * @Author: vincent
 * @Description:
 * @File:  client
 * @Version: 1.0.0
 * @Date: 2021/8/14 22:22
 */

package mock

import (
	"context"
	"go-examples/chat_demo_im/tcp"
	"time"

	kim "go-examples/chat_demo_im"
	"go-examples/chat_demo_im/logger"
	"go-examples/chat_demo_im/websocket"
	"net"

	"github.com/gobwas/ws/wsutil"

	"github.com/gobwas/ws"
)

// client demo
type ClientDemo struct {
}

func (c *ClientDemo) Start(userId, protocol, addr string) {
	var cli kim.Client

	// 1.初始化客户端
	if protocol == "ws" {
		// websocket
		// todo: name 和 options 应该可以配置
		cli = websocket.NewClient(userId, "ws client", websocket.ClientOptions{})
		// set dialer
		cli.SetDialer(&WebsocketDialer{})
	} else if protocol == "tcp" {
		cli = tcp.NewClient(userId, "tcp client", tcp.ClientOptions{})
		cli.SetDialer(&TcpDialer{})
	}

	// 2.建立连接
	err := cli.Connect(addr)
	if err != nil {
		logger.Error(err)
	}

	// 3.goroutine 发5次消息 内容是"hello"
	count := 5
	go func() {
		for i := 0; i < count; i++ {
			err := cli.Send([]byte("hello"))
			if err != nil {
				logger.Error(err)
				return
			}

			// sleep 10ms
			time.Sleep(time.Millisecond * 10)
		}
	}()

	// 4.接收消息
	recv := 0
	for {
		frame, err := cli.Read()
		if err != nil {
			logger.Info(err)
			break
		}
		if frame.GetOpCode() != kim.OpBinary {
			logger.Error("get msg not OpBinary")
			continue
		}

		recv++
		logger.Warnf("%s receive message [%s]", cli.ID(), frame.GetPayload())
		if recv == count {
			// 接收完消息
			break
		}
	}

	// 退出
	// todo: 这个不应该使用defer吗
	cli.Close()
}

type WebsocketDialer struct {
}

func (d *WebsocketDialer) DialAndHandshake(ctx kim.DialContext) (net.Conn, error) {
	// 1.调用ws.Dial拨号
	// todo: kim.DialContext感觉内容也没啥有用的, 叫context有点大
	conn, _, _, err := ws.Dial(context.TODO(), ctx.Address)
	if err != nil {
		return nil, err
	}

	// 2.发送用户认证信息，示例就是userId
	// 为啥是Binary呢，是因为第一条dial消息吗;我认为是任意，也可以用WriteClientText
	// 只要和server约定好就可以了
	// WriteClientBinary is the same as WriteClientMessage with ws.OpBinary.
	err = wsutil.WriteClientBinary(conn, []byte(ctx.Id))
	if err != nil {
		return nil, err
	}

	// 3.return conn
	return conn, nil
}

type TcpDialer struct {
}

func (d *TcpDialer) DialAndHandshake(ctx kim.DialContext) (net.Conn, error) {
	logger.Info("start dial:", ctx.Address)

	// 1.调用net.Dial拨号
	conn, err := net.DialTimeout("tcp", ctx.Address, ctx.Timeout)
	if err != nil {
		return nil, err
	}

	// 2.发送用户认证信息,示例就是userId
	err = tcp.WriteFrame(conn, kim.OpBinary, []byte(ctx.Id))
	if err != nil {
		return nil, err
	}

	// 3.return conn
	return conn, nil
}

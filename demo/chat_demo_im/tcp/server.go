/**
 * @Author: vincent
 * @Description:
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2021/8/13 15:41
 */

package tcp

import (
	"context"
	"errors"
	"fmt"
	kim "go-examples/demo/chat_demo_im"
	"go-examples/demo/chat_demo_im/logger"
	"go-examples/demo/chat_demo_im/naming"
	"net"
	"sync"
	"time"

	"github.com/segmentio/ksuid"
)

// ServerOptions
type ServerOptions struct {
	loginWait time.Duration //登录超时
	readWait  time.Duration //读超时
	writeWait time.Duration //写超时
}

// tcp implementation of the Server
// vincent: 注意下面都是小写的
type Server struct {
	listen string
	naming.ServiceRegistration
	kim.ChannelMap
	kim.Acceptor
	kim.MessageListener
	kim.StateListener
	once    sync.Once
	options ServerOptions
	quit    *kim.Event
}

func (s *Server) SetAcceptor(acceptor kim.Acceptor) {
	s.Acceptor = acceptor
}

func (s *Server) SetMessageListener(listener kim.MessageListener) {
	s.MessageListener = listener
}

func (s *Server) SetStateListener(listener kim.StateListener) {
	s.StateListener = listener
}

func (s *Server) SetChannelMap(channelMap kim.ChannelMap) {
	s.ChannelMap = channelMap
}

func (s *Server) SetReadWait(duration time.Duration) {
	s.options.readWait = duration
}

func (s *Server) Start() error {

	log := logger.WithFields(logger.Fields{
		"module": "tcp.server",
		"listen": s.listen,
		"id":     s.ServiceID(),
	})

	if s.StateListener == nil {
		return fmt.Errorf("server StateListener is null")
	}

	if s.Acceptor == nil {
		// todo: 不想使用原来的defaultAcceptor
		s.Acceptor = new(defaultAcceptor)
		//return fmt.Errorf("server Acceptor is null")
	}

	// listen
	listener, err := net.Listen("tcp", s.listen)
	if err != nil {
		return err
	}
	log.Info("server started")

	for {
		rawConn, err := listener.Accept()
		if err != nil {
			rawConn.Close()
			log.Warn(err)
			// todo：这个continue也很奇怪呢
			continue
		}

		go func(rawConn net.Conn) {
			conn := NewTcpConn(rawConn)

			// accept
			id, err := s.Accept(conn, s.options.loginWait)
			if err != nil {
				// accept fail, 关闭conn
				// todo: 是否需要发送OpClose
				_ = conn.WriteFrame(kim.OpClose, []byte(err.Error()))
				conn.Close()
				return
			}

			// 判断id是否存在
			if _, ok := s.Get(id); ok {
				// 存在
				log.Warnf("channel %s existed", id)
				_ = conn.WriteFrame(kim.OpClose, []byte("channelId is repeated"))
				conn.Close()
				return
			}

			// 添加channel
			channel := kim.NewChannel(id, conn)
			channel.SetReadWait(s.options.readWait)
			channel.SetWriteWait(s.options.writeWait)
			s.Add(channel)

			//readLoop
			log.Info("accept ", channel)
			err = channel.ReadLoop(s.MessageListener)
			if err != nil {
				log.Info(err)
			}

			// todo: 这个扫尾工作，感觉很乱呢
			// 删除channel
			s.Remove(channel.ID())
			// todo：断连具体做啥呢
			_ = s.Disconnect(channel.ID())
			// channel关闭
			channel.Close()
		}(rawConn)

		// 看下是否接到退出的通知
		select {
		case <-s.quit.Done():
			return fmt.Errorf("listen exited")
		default:
		}
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	log := logger.WithFields(logger.Fields{
		"module": "tcp.server",
		"id":     s.ServiceID(),
	})

	// do
	s.once.Do(func() {
		// defer
		defer func() {
			log.Infoln("shutdown")
		}()

		// close channels
		channels := s.ChannelMap.All()
		for _, ch := range channels {
			ch.Close()

			// todo: 总感觉这么写哪里不对
			select {
			case <-ctx.Done():
				// 收到结束通知
				return
			default:
				continue
			}
		}
	})

	return nil
}

func (s *Server) Push(channelId string, msg []byte) error {
	ch, ok := s.ChannelMap.Get(channelId)
	if !ok {
		return errors.New("channel no found")
	}
	return ch.Push(msg)
}

// NewServer
func NewServer(listen string, service naming.ServiceRegistration) kim.Server {
	return &Server{
		listen:              listen,
		ServiceRegistration: service,
		ChannelMap:          kim.NewChannelMap(100),
		//Acceptor:            nil,
		//MessageListener: nil,
		//StateListener: nil,
		//once:    sync.Once{},
		options: ServerOptions{
			loginWait: kim.DefaultLoginWait,
			readWait:  kim.DefaultReadWait,
			writeWait: time.Second * 10,
		},
		quit: kim.NewEvent(),
	}
}

// defaultAcceptor
type defaultAcceptor struct {
}

func (d *defaultAcceptor) Accept(conn kim.Conn, wait time.Duration) (string, error) {
	return ksuid.New().String(), nil
}

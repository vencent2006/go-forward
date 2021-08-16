/**
 * @Author: vincent
 * @Description:
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2021/8/14 23:23
 */

package websocket

import (
	"context"
	"fmt"
	kim "go-examples/chat_demo_im"
	"go-examples/chat_demo_im/logger"
	"go-examples/chat_demo_im/naming"
	"net/http"
	"sync"
	"time"

	"github.com/gobwas/ws"

	"github.com/segmentio/ksuid"
)

// server options
type ServerOptions struct {
	loginWait time.Duration // 登录超时
	readWait  time.Duration // 读超时
	writeWait time.Duration // 写超时
}

// Server is a websocket implement of the kim.Server
type Server struct {
	// control
	once sync.Once

	// session
	kim.ChannelMap

	// options
	listen string
	naming.ServiceRegistration
	kim.Acceptor
	kim.MessageListener
	kim.StateListener
	options ServerOptions
}

func (s *Server) SetAcceptor(acceptor kim.Acceptor) {
	s.Acceptor = acceptor
}

func (s *Server) SetMessageListener(listener kim.MessageListener) {
	s.MessageListener = listener
}

func (s *Server) SetStateListener(listerner kim.StateListener) {
	s.StateListener = listerner
}

func (s *Server) SetChannelMap(channelMap kim.ChannelMap) {
	s.ChannelMap = channelMap
}

func (s *Server) SetReadWait(duration time.Duration) {
	s.options.readWait = duration
}

func (s *Server) Start() error {
	// 1.http mux
	mux := http.NewServeMux()
	// logger
	log := logger.WithFields(logger.Fields{
		"module": "ws.server",
		"listen": s.listen,
		"id":     s.ServiceID(),
	})

	// 2.acceptor listener channelMap
	if s.Acceptor == nil {
		s.Acceptor = new(defaultAcceptor)
	}
	if s.StateListener == nil {
		return fmt.Errorf("StateListener is nil")
	}
	if s.ChannelMap == nil {
		s.ChannelMap = kim.NewChannelMap(100)
	}

	// 3.http mux handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 1.upgrade protocol
		rawConn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			resp(w, http.StatusBadRequest, err.Error())
			return
		}

		// 2.包装conn
		conn := NewWsConn(rawConn)

		// 3.accept
		// todo: websocket 还有accept吗
		id, err := s.Accept(conn, s.options.loginWait)
		if err != nil {
			// todo: accept fail, 还要通知close
			_ = conn.WriteFrame(kim.OpClose, []byte(err.Error()))
			conn.Close()
			return
		}

		// 4.add channel
		if _, ok := s.Get(id); ok {
			// todo: id已经连接了，当前是不kickoff的，未来要kickoff吗
			log.Warnf("channel %s existed", id)
			_ = conn.WriteFrame(kim.OpClose, []byte("channelId is repeated"))
			conn.Close()
			return
		}

		channel := kim.NewChannel(id, conn)
		channel.SetWriteWait(s.options.writeWait)
		channel.SetReadWait(s.options.readWait)
		s.Add(channel)

		// 5.起个goroutine处理消息
		go func(ch kim.Channel) {
			// 6.read loop
			err := ch.ReadLoop(s.MessageListener)
			if err != nil {
				log.Info(err)
			}

			// 7.remove channel
			s.Remove(ch.ID())
			err = s.Disconnect(ch.ID())
			if err != nil {
				log.Warnf("channel(%s) disconnect err(%s)", ch.ID(), err)
			}
			ch.Close()
		}(channel)

	})
	log.Infof("%s:%s started", s.ServiceID(), s.ServiceName())

	// 7:Listen and Serve
	return http.ListenAndServe(s.listen, mux)
}

func (s *Server) Shutdown(ctx context.Context) error {
	log := logger.WithFields(logger.Fields{
		"module": "ws.server",
		"id":     s.ServiceID(),
	})

	s.once.Do(func() {
		// defer log
		defer func() {
			log.Infoln("shutdown")
		}()

		// close channels(channelMap)
		channels := s.ChannelMap.All()
		for _, ch := range channels {
			// close
			ch.Close()

			// 检查是否收到结束通知
			select {
			case <-ctx.Done():
				log.Infoln("receive ctx.Done")
				return
			default:
				continue
			}
		}
	})
	return nil
}

func (s *Server) Push(channelId string, data []byte) error {
	// 检查channel是否存在
	ch, ok := s.ChannelMap.Get(channelId)
	if !ok {
		return fmt.Errorf("channel(%s) not found", channelId)
	}
	return ch.Push(data)
}

func resp(w http.ResponseWriter, code int, body string) {
	// set header
	w.WriteHeader(code)

	// write
	if body != "" {
		_, _ = w.Write([]byte(body))
	}

	logger.Infof("response with code:%d, %s", code, body)
}

// NewServer
func NewServer(listen string, service naming.ServiceRegistration) kim.Server {
	return &Server{
		listen:              listen,
		ServiceRegistration: service,
		options: ServerOptions{
			loginWait: kim.DefaultLoginWait,
			readWait:  kim.DefaultReadWait,
			writeWait: time.Second * 10,
		},
	}
}

// defaultAcceptor
type defaultAcceptor struct {
}

func (a *defaultAcceptor) Accept(conn kim.Conn, wait time.Duration) (string, error) {
	return ksuid.New().String(), nil
}

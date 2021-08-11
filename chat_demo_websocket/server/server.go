/**
 * @Author: vincent
 * @Description:
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2021/8/10 09:32
 */

package server

import (
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/gobwas/ws/wsutil"

	"github.com/gobwas/ws"
	"github.com/sirupsen/logrus"
)

const (
	MAX_USER_NUM = 100 // 保持100个用户的session
)

// server is websocket server
type Server struct {
	id   string
	addr string
	// 锁
	sync.Mutex
	// 只执行一次，用来关闭
	once sync.Once
	// session
	users map[string]net.Conn
}

func NewServer(id, addr string) *Server {
	return newServer(id, addr, MAX_USER_NUM)
}

func newServer(id string, addr string, num int) *Server {
	return &Server{
		id:    id,
		addr:  addr,
		users: make(map[string]net.Conn, num),
	}
}

func (s *Server) addUser(user string, conn net.Conn) (net.Conn, bool) {
	s.Lock()
	defer s.Unlock()

	old, ok := s.users[user] // 返回旧的连接
	s.users[user] = conn     // 更新为新的连接

	return old, ok
}

func (s *Server) delUser(user string) {
	s.Lock()
	defer s.Unlock()

	delete(s.users, user)
}

func (s *Server) readLoop(userId string, conn net.Conn) error {
	for {
		frame, err := ws.ReadFrame(conn)
		if err != nil {
			return fmt.Errorf("%s ReadFrame fail, %v", userId, err)
			return err
		}

		if frame.Header.OpCode == ws.OpClose {
			// 对端关闭
			return fmt.Errorf("%s remote side close the conn", userId)
		}

		if frame.Header.OpCode == ws.OpPing {
			logrus.Info("recv a ping; resp with a pong")
			// f := ws.NewPongFrame(nil)
			// _ = ws.WriteFrame(conn, f)
			// 上面两句，等价于 wsutil.WriteServerMessage(conn, ws.OpPong, nil)
			_ = wsutil.WriteServerMessage(conn, ws.OpPong, nil)
			continue
		}

		// use mask to decode
		if frame.Header.Masked {
			ws.Cipher(frame.Payload, frame.Header.Mask, 0)
		}

		// 接收文本信息
		if frame.Header.OpCode == ws.OpText {
			go s.handle(userId, string(frame.Payload))
		}
	}
}

func (s *Server) handle(userId string, message string) {
	logrus.Infof("recv msg %s, from %s\n", message, userId)

	s.Lock()
	defer s.Unlock()

	// 发广播消息
	broadcastMsg := fmt.Sprintf("%s --- From %s", message, userId)
	for user, conn := range s.users {
		if user == userId {
			continue
		}
		logrus.Infof("send to %s: %s\n", user, broadcastMsg)
		err := s.writeText(conn, broadcastMsg)
		if err != nil {
			logrus.Errorf("write to %s failed, err is %v", user, err)
		}
	}
}

func (s *Server) writeText(conn net.Conn, message string) error {
	frame := ws.NewTextFrame([]byte(message))
	return ws.WriteFrame(conn, frame)
}

/*
**************** public function
 */

func (s *Server) Shutdown() {
	s.once.Do(func() {
		s.Lock()
		defer s.Unlock()

		// 关闭所有session
		for i := range s.users {
			s.users[i].Close()
		}
	})
}

func (s *Server) Start() error {
	log := logrus.WithFields(logrus.Fields{
		"module": "Server",
		"id":     s.id,
		"addr":   s.addr,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//1.step update http to websocket
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Errorf("update websocket fail, err:%v", err)
			conn.Close()
			return
		}
		//2.get user id
		userId := r.URL.Query().Get("user")
		if userId == "" {
			log.Errorln("get user id fails")
			conn.Close()
			return
		}
		//3.add user to session
		old, ok := s.addUser(userId, conn)
		if ok {
			// 已经存在连接, 踢掉old连接
			old.Close()
			log.Infof("already connect, kickoff old(%s)", userId)
		}
		log.Infof("user(%s) come in.", userId)

		//4.go func process
		go func(userId string, conn net.Conn) {
			// 4.读取消息
			err := s.readLoop(userId, conn)
			if err != nil {
				log.Errorln(err)
			}
			// 5.断开连接，删除用户
			conn.Close()
			s.delUser(userId)

			log.Infof("connection of %s closed", userId)
		}(userId, conn)

	})

	log.Infoln("Server started")
	return http.ListenAndServe(s.addr, mux)
}

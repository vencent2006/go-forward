/**
 * @Author: vincent
 * @Description:
 * @File:  util
 * @Version: 1.0.0
 * @Date: 2022/3/1 17:16
 */

package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/gobwas/ws"
)

// MAX_CONN_NUM 最大连接数量
const MAX_CONN_NUM = 10

// Server 核心结构
type Server struct {
	once sync.Once // 保证shutdown操作，只执行1次

	id      string // server id标识
	address string // 监听地址

	sync.Mutex                     // lock
	users      map[string]net.Conn // user connection session
}

// NewServer 构造函数
func NewServer(id, address string) *Server {
	return newServer(id, address)
}

// newServer 私有构造方法
func newServer(id, address string) *Server {
	return &Server{
		id:      id,
		address: address,
		users:   make(map[string]net.Conn, MAX_CONN_NUM),
	}
}

func (s *Server) Start() error {
	log := logrus.WithFields(logrus.Fields{
		"module": "Server",
		"listen": s.address,
		"id":     s.id,
	})
	// mux
	mux := http.NewServeMux()

	// mux handleFunc
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// upgrade websocket
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Errorf("upgrade http fail: %v", err)
			conn.Close()
			return
		}

		// get user
		user := r.URL.Query().Get("user")
		if user == "" {
			log.Errorf("get user fail")
			conn.Close()
			return
		}

		// add user
		old, isExist := s.addUser(user, conn)
		if isExist {
			// old存在，要close掉
			old.Close()
		}

		log.Infof("user(%s) come in", user)

		// go readLoop
		go func(user string, conn net.Conn) {
			err = s.readLoop(user, conn)
			if err != nil {
				logrus.Infof("user(%s) close connection, err:%v", user, err)
			}
			conn.Close()
			s.delUser(user)
			log.Infof("connection of %s closed", user)
		}(user, conn)
	})

	// listenAndServ
	log.Info("Server started...")
	return http.ListenAndServe(s.address, mux)
}

func (s *Server) addUser(user string, conn net.Conn) (net.Conn, bool) {
	s.Lock()
	defer s.Unlock()
	old, ok := s.users[user]
	s.users[user] = conn
	return old, ok
}

func (s *Server) delUser(user string) {
	s.Lock()
	defer s.Unlock()
	delete(s.users, user)
}

// Shutdown 关闭连接
func (s *Server) Shutdown() {
	s.once.Do(func() {
		s.Lock()
		defer s.Unlock()
		for _, conn := range s.users {
			conn.Close()
		}
	})
}

func (s *Server) readLoop(user string, conn net.Conn) error {
	for {
		// read frame
		frame, err := ws.ReadFrame(conn)
		if err != nil {
			return err
		}

		// if closed
		if frame.Header.OpCode == ws.OpClose {
			// 对端关闭
			return errors.New("remote client close conn")
		}

		// masked
		if frame.Header.Masked {
			ws.Cipher(frame.Payload, frame.Header.Mask, 0)
		}

		// text
		if frame.Header.OpCode == ws.OpText {
			go s.handle(user, string(frame.Payload))
		}
	}
}

func (s *Server) handle(user string, text string) {
	// lock
	s.Lock()
	defer s.Unlock()

	// broadcast
	broadcastMsg := fmt.Sprintf("receive: %s -- FROM %s", text, user)
	for u, conn := range s.users {
		if u == user {
			// 自己不发消息
			continue
		}

		err := s.writeText(conn, broadcastMsg)
		if err != nil {
			logrus.Errorf("write to %s fail, err:%v", u, err)
		}
	}
}

// 写websocket的text（文本）
func (s *Server) writeText(conn net.Conn, message string) error {
	f := ws.NewTextFrame([]byte(message))
	return ws.WriteFrame(conn, f)
}

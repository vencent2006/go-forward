/**
 * @Author: vincent
 * @Description:
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2022/3/1 19:46
 */

package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/gobwas/ws"

	"github.com/sirupsen/logrus"
)

// MAX_CONN_NUM 最大连接数
const MAX_CONN_NUM = 10

type Server struct {
	once    sync.Once
	address string
	id      string
	sync.Mutex
	users map[string]net.Conn
}

func NewServer(id, address string) *Server {
	return newServer(id, address)
}

func newServer(id, address string) *Server {
	return &Server{
		address: address,
		id:      id,
		users:   make(map[string]net.Conn, 10),
	}
}

func (s *Server) Start() error {
	// log
	log := logrus.WithFields(logrus.Fields{
		"module":  "server",
		"address": s.address,
		"id":      s.id,
	})

	// mux
	mux := http.NewServeMux()

	// mux handleFunc
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http upgrade to websocket
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Errorf("UpgradeHTTP fail, err:%v", err)
			conn.Close()
			return
		}

		// get user
		user := r.URL.Query().Get("user")
		if user == "" {
			// get user fail
			log.Errorf("get user fail")
			conn.Close()
			return
		}

		// add user and conn
		old, ok := s.addUser(user, conn)
		if ok {
			log.Infof("old(%s) ready to close", user)
			old.Close()
			return
		}

		// info
		log.Infof("user(%s) come in.", user)

		// go readLoop
		go func(user string, conn net.Conn) {
			err = s.readLoop(user, conn)
			if err != nil {
				log.Infof("recv close from user(%s), err:%v", user, err)
			}
			conn.Close()
			s.delUser(user)

		}(user, conn)

	})

	log.Info("server start...")

	// listenAndServe
	return http.ListenAndServe(s.address, mux)
}

func (s *Server) Shutdown() {
	s.once.Do(func() {
		// lock
		s.Lock()
		defer s.Unlock()

		// close conn
		for _, conn := range s.users {
			conn.Close()
		}
	})
}

// private function
func (s *Server) addUser(user string, conn net.Conn) (net.Conn, bool) {
	// lock
	s.Lock()
	defer s.Unlock()

	old, ok := s.users[user]
	s.users[user] = conn

	return old, ok
}

func (s *Server) delUser(user string) {
	// mutex
	s.Lock()
	defer s.Unlock()

	delete(s.users, user)
}

func (s *Server) readLoop(user string, conn net.Conn) error {
	for {
		// read frame
		frame, err := ws.ReadFrame(conn)
		if err != nil {
			return err
		}

		// is closed?
		if frame.Header.OpCode == ws.OpClose {
			return errors.New("remote client close connection")
		}

		// is masked?
		if frame.Header.Masked {
			ws.Cipher(frame.Payload, frame.Header.Mask, 0)
		}

		// text
		if frame.Header.OpCode == ws.OpText {
			go s.handle(user, string(frame.Payload))
		}
	}
}

func (s *Server) handle(user, text string) {
	//mutex
	s.Lock()
	defer s.Unlock()

	// 构造广播消息
	broadcastMsg := fmt.Sprintf("recv: %s from %s", text, user)

	// send message
	for u, conn := range s.users {
		// 不给自己发消息
		if u == user {
			continue
		}

		frame := ws.NewTextFrame([]byte(broadcastMsg))
		err := ws.WriteFrame(conn, frame)
		if err != nil {
			fmt.Printf("WriteFrame to %s fail, err:%v", u, err)
		}
	}
}

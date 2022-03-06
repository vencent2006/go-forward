/**
 * @Author: vincent
 * @Description:
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2022/3/1 23:12
 */

package server

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gobwas/ws/wsutil"

	"github.com/gobwas/ws"

	"github.com/sirupsen/logrus"
)

const MAX_CONN_NUM = 10
const MAX_WRITE_WAIT = time.Second * 10
const MAX_READ_WAIT = time.Minute * 2

func runServer(ctx context.Context, opts *options, version string) error {
	s := newServer(opts.id, opts.listen)
	defer s.shutdown()
	return s.start()
}

type serverOptions struct {
	// 写等待
	writeWait time.Duration
	// 读等待
	readWait time.Duration
}

type server struct {
	id     string
	listen string

	// control
	sync.Once
	sync.Mutex
	opts serverOptions

	// sessions
	users map[string]net.Conn
}

func newServer(id, listen string) *server {
	return &server{
		id:     id,
		listen: listen,
		opts: serverOptions{
			writeWait: MAX_WRITE_WAIT,
			readWait:  MAX_READ_WAIT,
		},
		users: make(map[string]net.Conn, MAX_CONN_NUM),
	}
}

func (s *server) shutdown() {

}

func (s *server) start() error {
	// set log
	logrus.WithFields(logrus.Fields{
		"module": "server",
		"id":     s.id,
		"listen": s.listen,
	})

	// mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get user
		user := r.URL.Query().Get("user")
		if user == "" {
			logrus.Errorf("get user fail")
			return
		}

		// upgrade http
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			logrus.Errorf("upgrade http fail, err: %v", err)
			return
		}

		// add session
		old, ok := s.addUser(user, conn)
		if ok {
			old.Close()
			logrus.Warnf("close old(%s) connection(%s)", user, old.RemoteAddr())
		}

		// read loop
		go func(user string, conn net.Conn) {

			// readLoop
			err = s.readLoop(user, conn)
			if err != nil {
				logrus.Warnf("readLoop exit - %v", err)
			}
			conn.Close()

			// delete user
			s.delUser(user)

			// log
			logrus.Infof("connection of %s closed", user)
		}(user, conn)

		// write loop
	})

	logrus.Info("server start...")
	return http.ListenAndServe(s.listen, mux)
}

func (s *server) readLoop(user string, conn net.Conn) error {
	for {
		frame, err := ws.ReadFrame(conn)
		if err != nil {
			return err
		}

		switch frame.Header.OpCode {
		case ws.OpClose: // close
			logrus.Info("r")
		case ws.OpPing: // ping
			wsutil.WriteServerMessage(conn, ws.OpPong, nil)
			logrus.Infof("write pong to %s", user)
		case ws.OpText: // text
			if frame.Header.Masked {
				ws.Cipher(frame.Payload, frame.Header.Mask, 0)
			}
			go s.handleText(user, string(frame.Payload))
		case ws.OpBinary: // binary
			if frame.Header.Masked {
				ws.Cipher(frame.Payload, frame.Header.Mask, 0)
			}
			go s.handleBinary(user, frame.Payload)
		}
	}
}

func (s *server) addUser(user string, conn net.Conn) (net.Conn, bool) {
	s.Lock()
	defer s.Unlock()
	old, ok := s.users[user]
	s.users[user] = conn
	return old, ok
}

func (s *server) delUser(user string) error {
	s.Lock()
	defer s.Unlock()
	delete(s.users, user)
	return nil
}

func (s *server) handleText(user string, text string) {
	logrus.Infof("receive msg_text(%s) from user(%s)", text, user)
	// lock mutex
	s.Lock()
	defer s.Unlock()

	broadcastMsg := fmt.Sprintf("%s ---- from %s", text, user)
	for u, conn := range s.users {
		if u == user {
			// 不发给自己
			continue
		}

		// log
		logrus.Infof("send to %s , %s", u, broadcastMsg)

		// write text
		err := s.writeText(conn, broadcastMsg)
		if err != nil {
			logrus.Errorf("write broadcastMsg(%s) to %s fail, err:%v", broadcastMsg, u, err)
		}
	}
}

const (
	commandPing = 100
	commandPong = 101
)

func (s *server) handleBinary(user string, message []byte) {
	// log
	logrus.Infof("receive msg_binary(%v) from user(%s)", message, user)

	// lock mutex
	s.Lock()
	defer s.Unlock()

	// unpack message
	// handle ping request
	i := 0
	command := binary.BigEndian.Uint16(message[i : i+2])
	i += 2
	payloadLen := binary.BigEndian.Uint32(message[i : i+4])

	//log
	logrus.Infof("receive binary msg: command=%v, payloadLen=%v", command, payloadLen)

	//根据command进行回复
	if command == commandPing {
		// 收到receive commandPing
		conn := s.users[user]
		// return conn
		err := wsutil.WriteServerBinary(conn, []byte{commandPong, 0, 0, 0, 0})
		if err != nil {
			logrus.Errorf("write to user(%s) fail, err:%v", user, err)
			return
		}
	}

}

func (s *server) writeText(conn net.Conn, text string) error {
	f := ws.NewTextFrame([]byte(text))
	err := conn.SetWriteDeadline(time.Now().Add(s.opts.writeWait))
	if err != nil {
		return err
	}
	return ws.WriteFrame(conn, f)
}

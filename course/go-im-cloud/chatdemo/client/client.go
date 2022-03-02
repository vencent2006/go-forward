/**
 * @Author: vincent
 * @Description:
 * @File:  client
 * @Version: 1.0.0
 * @Date: 2022/3/1 23:13
 */

package client

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/gobwas/ws/wsutil"

	"github.com/gobwas/ws"

	"github.com/sirupsen/logrus"
)

func runClient(ctx context.Context, opts *Options) error {
	logrus.WithFields(logrus.Fields{
		"module": "client",
		"server": opts.address,
		"user":   opts.user,
	})

	// connect
	url := fmt.Sprintf("%s?user=%s", opts.address, opts.user)
	logrus.Info("connect to ", url)
	h, err := connect(url)
	if err != nil {
		return err
	}
	<-h.close
	logrus.Info("connection closed")

	return nil
}

type handler struct {
	conn      net.Conn
	close     chan struct{}
	recv      chan []byte
	heartbeat time.Duration
}

func newHandler(addr string) (*handler, error) {
	// check addr
	_, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	// dial
	conn, _, _, err := ws.Dial(context.Background(), addr)
	if err != nil {
		return nil, err
	}

	// ret
	return &handler{
		conn:      conn,
		close:     make(chan struct{}, 1),
		recv:      make(chan []byte, 10),
		heartbeat: time.Second * 50,
	}, nil
}

func connect(address string) (*handler, error) {
	// check address
	h, err := newHandler(address)
	if err != nil {
		return nil, err
	}

	// readLoop
	go func() {
		err := h.readLoop()
		if err != nil {
			logrus.Warnf("readLoop end, err:%v", err)
		}

		// notify close
		h.close <- struct{}{}
	}()

	// heartbeat
	go func() {
		err := h.heartbeatLoop()
		if err != nil {
			logrus.Infof("heartbeat loop, err:%v", err)
		}
	}()

	// receive msg
	go func() {
		for msg := range h.recv {
			logrus.Info("receive msg: ", msg)
		}
	}()

	// send msg
	go func() {
		ticker := time.NewTicker(time.Second * 6)
		msg := "hello"
		for range ticker.C {
			h.setWriteDeadline(time.Second * 10)
			err := wsutil.WriteClientText(h.conn, []byte(msg))
			if err != nil {
				logrus.Error(err)
			}
		}
	}()

	return h, nil
}
func (h *handler) setWriteDeadline(d time.Duration) error {
	return h.conn.SetDeadline(time.Now().Add(d))
}

func (h *handler) readLoop() error {
	// log
	logrus.Info("enter readLoop")

	// set deadline
	err := h.setWriteDeadline(h.heartbeat * 3)
	if err != nil {
		return err
	}

	// read frame
	for {
		frame, err := ws.ReadFrame(h.conn)
		if err != nil {
			return err
		}

		switch frame.Header.OpCode {
		case ws.OpPong:
			logrus.Info("receive pong")
			h.setWriteDeadline(h.heartbeat * 3)
		case ws.OpClose:
			logrus.Info("remote close")
			return errors.New("remote close")
		case ws.OpText:
			// todo: 这个会阻塞
			h.recv <- frame.Payload
		}
	}
}

func (h *handler) heartbeatLoop() error {
	logrus.Info("enter heartbeat loop")

	ticker := time.NewTicker(h.heartbeat)
	for range ticker.C {
		logrus.Info("send ping")
		err := h.setWriteDeadline(time.Second * 10)
		if err != nil {
			return err
		}
		err = wsutil.WriteClientMessage(h.conn, ws.OpPing, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

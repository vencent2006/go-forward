/**
 * @Author: vincent
 * @Description:
 * @File:  client
 * @Version: 1.0.0
 * @Date: 2021/8/14 23:22
 */

package websocket

import (
	"errors"
	"fmt"
	kim "go-examples/demo/chat_demo_im"
	"go-examples/demo/chat_demo_im/logger"
	"net"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type ClientOptions struct {
	Heartbeat time.Duration
	ReadWait  time.Duration
	WriteWait time.Duration
}

// Client is a websocket implement of the terminal
type Client struct {
	// control
	sync.Mutex
	once sync.Once

	// dialer
	kim.Dialer
	dc *kim.DialContext

	// attr
	id      string
	name    string
	state   int32
	options ClientOptions
	conn    net.Conn
}

// NewClient
func NewClient(id, name string, opts ClientOptions) kim.Client {
	// 1.替换默认值
	if opts.WriteWait == 0 {
		opts.WriteWait = kim.DefaultWriteWait
	}
	if opts.ReadWait == 0 {
		opts.ReadWait = kim.DefaultReadWait
	}

	// 2.new Client
	cli := &Client{
		id:      id,
		name:    name,
		options: opts,
	}

	// return
	return cli
}

func (c *Client) SetDialer(dialer kim.Dialer) {
	c.Dialer = dialer
}

func (c *Client) ID() string {
	return c.id
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) Connect(addr string) error {
	// 1.check addr
	_, err := url.Parse(addr)
	if err != nil {
		return err
	}

	// 2.check state
	if !atomic.CompareAndSwapInt32(&c.state, 0, 1) {
		return fmt.Errorf("client has connected")
	}

	// 3.dial
	conn, err := c.Dialer.DialAndHandshake(kim.DialContext{
		Id:      c.id,
		Name:    c.name,
		Address: addr,
		Timeout: kim.DefaultLoginWait,
	})
	if err != nil {
		atomic.CompareAndSwapInt32(&c.state, 1, 0)
		return err
	}
	// todo: 存在err = nil，但是conn = nil的情况，是什么情况呢
	if conn == nil {
		return fmt.Errorf("conn is nil")
	}

	// 4.set conn
	c.conn = conn

	// 5.heartbeat
	if c.options.Heartbeat > 0 {
		go func() {
			//err := c.h
		}()
	}

	return nil
}

func (c *Client) Send(payload []byte) error {
	// 1.check state: require state = 0
	if atomic.LoadInt32(&c.state) == 0 {
		return fmt.Errorf("connection is nil")
	}

	// 2.lock
	c.Lock()
	defer c.Unlock()

	// 3.set write deadline
	err := c.conn.SetWriteDeadline(time.Now().Add(c.options.WriteWait))
	if err != nil {
		return err
	}

	// 4.write message, 客户端信息需要使用mask
	return wsutil.WriteClientMessage(c.conn, ws.OpBinary, payload)
}

func (c *Client) Read() (kim.Frame, error) {
	// 1.check conn
	if c.conn == nil {
		return nil, errors.New("connection is nil")
	}
	// 2.set deadline if heartbeat
	if c.options.Heartbeat > 0 {
		_ = c.conn.SetReadDeadline(time.Now().Add(c.options.ReadWait))
	}
	// 3.read frame
	frame, err := ws.ReadFrame(c.conn)
	if err != nil {
		return nil, err
	}

	// 4.OpClose?
	if frame.Header.OpCode == ws.OpClose {
		// close from server
		return nil, fmt.Errorf("remote side close the channel %s", c.id)
	}

	// 4.return frame
	return &Frame{frame}, nil
}

func (c *Client) Close() {
	c.once.Do(func() {
		// check conn
		if c.conn == nil {
			// conn为空
			return
		}

		// 优雅关闭 graceful close connection
		_ = wsutil.WriteClientMessage(c.conn, ws.OpClose, nil)

		// close
		c.conn.Close()

		// set state: 1 -> 0
		atomic.CompareAndSwapInt32(&c.state, 1, 0)
	})
}

func (c *Client) heartbeatLoop(conn net.Conn) error {
	tick := time.NewTicker(c.options.Heartbeat)
	for range tick.C {
		// 发送一个ping的心跳包给server
		if err := c.ping(conn); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ping(conn net.Conn) error {
	// 1.lock and unlock
	// todo：这里要加互斥锁，应该是writeFrame本身需要，考虑一下都哪些地方需要lock
	c.Lock()
	defer c.Unlock()

	// 2.write deadline
	// todo: 为什么不是wait，而是deadline呢
	err := conn.SetWriteDeadline(time.Now().Add(c.options.WriteWait))
	if err != nil {
		return err
	}

	// 3.log trace
	logger.Tracef("%s send ping to server", c.id)

	// 4.write client message
	return wsutil.WriteClientMessage(conn, ws.OpPing, nil)
}

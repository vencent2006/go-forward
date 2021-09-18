/**
 * @Author: vincent
 * @Description:
 * @File:  client
 * @Version: 1.0.0
 * @Date: 2021/8/13 15:41
 */

package tcp

import (
	"errors"
	"fmt"
	kim "go-examples/demo/chat_demo_im"
	"go-examples/demo/chat_demo_im/logger"
	"net/url"
	"sync"
	"sync/atomic"
	"time"
)

// client 选项
type ClientOptions struct {
	Heartbeat time.Duration // 心跳
	ReadWait  time.Duration // 读超时
	WriteWait time.Duration // 写超时
}

// client is a tcp implement of the terminal
type Client struct {
	// 控制信息
	sync.Mutex           // 防止并发进入临界区
	kim.Dialer           // 拨号器
	once       sync.Once //只执行一次

	// 属性信息
	id      string
	name    string
	conn    kim.Conn
	state   int32 // 应该是: 0是未连接，1是已连接
	options ClientOptions
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
	// 1.parse url
	_, err := url.Parse(addr)
	if err != nil {
		return err
	}

	// todo: 为什么要是一个cas呢
	// 2.设置state
	// 这里是一个cas原子操作，对比并设置值，是并发安全的
	if !atomic.CompareAndSwapInt32(&c.state, 0, 1) {
		return fmt.Errorf("client has connected")
	}

	// 3.conn
	rawConn, err := c.Dialer.DialAndHandshake(kim.DialContext{
		Id:      c.id,
		Name:    c.name,
		Address: addr,
		Timeout: kim.DefaultLoginWait,
	})
	if err != nil {
		atomic.CompareAndSwapInt32(&c.state, 1, 0)
		return err
	}
	if rawConn == nil {
		return fmt.Errorf("conn is nil")
	}
	c.conn = NewTcpConn(rawConn) // 升级为tcpConn

	// 4.heartbeat
	if c.options.Heartbeat > 0 {
		//如果设置了heartbeat，进行heartbeatLoop
		go func() {
			err := c.heartbeatLoop()
			if err != nil {
				logger.WithField("module", "tcp.client").Warn("heartbeatLoop stopped - ", err)
			}
		}()
	}

	return nil
}

func (c *Client) Send(payload []byte) error {
	// 1.检查state
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

	// 4.write frame
	return c.conn.WriteFrame(kim.OpBinary, payload)
}

func (c *Client) Read() (kim.Frame, error) {
	// 1.检查conn
	// todo: 为啥send检查state不检查conn,而Read要检查conn不检查state,有啥规律吗
	if c.conn == nil {
		return nil, errors.New("connection is nil")
	}

	// 2.查看heartbeat
	if c.options.Heartbeat > 0 {
		// todo: 有了心跳，为啥设置为ReadWait，而不是Heartbeat呢（Heartbeat只用来设置tick吧）
		_ = c.conn.SetReadDeadline(time.Now().Add(c.options.ReadWait))
	}

	// 3.read frame
	frame, err := c.conn.ReadFrame()
	if err != nil {
		return nil, err
	}

	// 4.跟进opCode进行判断
	if frame.GetOpCode() == kim.OpClose {
		// todo: 判断是对端close，不设置state吗
		return nil, errors.New("remote side close the channel")
	}

	return frame, nil
}

func (c *Client) Close() {
	// 只执行1次
	c.once.Do(func() {
		// todo: 不判断状态state吗
		if c.conn == nil {
			return
		}

		// gracefully close connection，对对端发close
		_ = WriteFrame(c.conn, kim.OpClose, nil)

		// conn close
		c.conn.Close()

		//设置状态
		atomic.CompareAndSwapInt32(&c.state, 1, 0)
	})
}

func NewClient(id string, name string, opts ClientOptions) kim.Client {

	// options
	if opts.WriteWait == 0 {
		opts.WriteWait = kim.DefaultWriteWait
	}
	if opts.ReadWait == 0 {
		opts.ReadWait = kim.DefaultReadWait
	}

	cli := &Client{
		id:      id,
		name:    name,
		options: opts,
	}
	return cli
}

func (c *Client) heartbeatLoop() error {
	tick := time.NewTicker(c.options.Heartbeat)
	for range tick.C {
		// 发送一个ping的心跳给server
		err := c.ping()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ping() error {
	logger.WithField("module", "tcp.client").Tracef("%s send ping to server", c.id)
	// todo: 每次都要SetWriteDeadline,真的很烦;
	// 为啥不能直接调用Send呢，或者把SetWriteDeadline防止在writeFrame中
	err := c.conn.SetWriteDeadline(time.Now().Add(c.options.WriteWait))
	if err != nil {
		return err
	}
	return c.conn.WriteFrame(kim.OpPing, nil)
}

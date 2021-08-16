/**
 * @Author: vincent
 * @Description:
 * @File:  client
 * @Version: 1.0.0
 * @Date: 2021/8/12 16:50
 */

package kim

import (
	"net"
	"time"
)

type Client interface {
	// setter
	SetDialer(Dialer)

	// getter
	ID() string
	Name() string

	// action
	Connect(addr string) error
	Send(payload []byte) error
	Read() (Frame, error)
	Close()
}

// 拨号器
type Dialer interface {
	DialAndHandshake(DialContext) (net.Conn, error)
}

type DialContext struct {
	Id      string
	Name    string
	Address string
	Timeout time.Duration
}

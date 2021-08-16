/**
 * @Author: vincent
 * @Description:
 * @File:  channel
 * @Version: 1.0.0
 * @Date: 2021/8/12 16:35
 */

package kim

import (
	"net"
	"time"
)

// Agent is a interface of client side
type Agent interface {
	ID() string
	Push(payload []byte) error
}

type Conn interface {
	net.Conn
	ReadFrame() (Frame, error)
	WriteFrame(code OpCode, payload []byte) error
	Flush() error
}

type Channel interface {
	Conn
	Agent

	// operation
	Close() error
	ReadLoop(lis MessageListener) error

	// set wait
	SetWriteWait(duration time.Duration)
	SetReadWait(duration time.Duration)
}

/**
 * @Author: vincent
 * @Description:
 * @File:  connection
 * @Version: 1.0.0
 * @Date: 2021/8/13 15:41
 */

package tcp

import (
	kim "go-examples/demo/chat_demo_im"
	"go-examples/demo/chat_demo_im/wire/endian"
	"io"
	"net"
)

// tcp frame
type Frame struct {
	OpCode  kim.OpCode
	Payload []byte
}

func (f *Frame) SetOpCode(code kim.OpCode) {
	f.OpCode = code
}

func (f *Frame) GetOpCode() kim.OpCode {
	return f.OpCode
}

func (f *Frame) SetPayload(payload []byte) {
	f.Payload = payload
}

func (f *Frame) GetPayload() []byte {
	return f.Payload
}

// tcp connection
type TcpConn struct {
	net.Conn
}

func NewTcpConn(conn net.Conn) *TcpConn {
	return &TcpConn{conn}
}

func (c *TcpConn) ReadFrame() (kim.Frame, error) {
	// read opCode
	opCode, err := endian.ReadUint8(c.Conn)
	if err != nil {
		return nil, err
	}

	// read payload
	payload, err := endian.ReadBytes(c.Conn)
	if err != nil {
		return nil, err
	}

	// return frame
	return &Frame{
		OpCode:  kim.OpCode(opCode),
		Payload: payload,
	}, nil
}

func (c *TcpConn) WriteFrame(code kim.OpCode, payload []byte) error {
	return WriteFrame(c.Conn, code, payload)
}

func (c *TcpConn) Flush() error {
	// todo: Flush待实现
	return nil
}

func WriteFrame(w io.Writer, code kim.OpCode, payload []byte) error {
	if err := endian.WriteUint8(w, uint8(code)); err != nil {
		return err
	}
	if err := endian.WriteBytes(w, payload); err != nil {
		return err
	}
	return nil
}

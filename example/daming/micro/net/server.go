package net

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

// 长度字段使用的字节数量
const numOfLengthBytes = 8

func Serve(network, addr string) error {
	listener, err := net.Listen(network, addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			// 一般是端口被占用的错误
			return err
		}

		go func() {
			if err := handleConn(conn); err != nil {
				_ = conn.Close()
			}
		}()
	}
}

func handleConn(conn net.Conn) error {
	for {
		bs := make([]byte, 8)
		_, err := conn.Read(bs)
		if err == net.ErrClosed || err == io.EOF || err == io.ErrUnexpectedEOF {
			// 这种是明确关闭的连接
			return err
		}
		if err != nil { // 这种是可以挽救的错误
			continue
		}
		//if n != 8 {
		//	return errors.New("micro: 没读够数据")
		//}

		res := handleMsg(bs)
		_, err = conn.Write(res)
		if err == net.ErrClosed || err == io.EOF || err == io.ErrUnexpectedEOF {
			// 这种是明确关闭的连接
			return err
		}
		if err != nil { // 这种是可以挽救的错误
			continue
		}
		//if n != len(res) {
		//	return errors.New("micro: 没写完数据")
		//}

	}
}

// 老师建议错误了就返回, 简单粗暴，不容易出bug
func handleConnV1(conn net.Conn) error {
	for {
		bs := make([]byte, 8)
		n, err := conn.Read(bs)
		if err != nil { // 老师建议错误了就返回, 简单粗暴，不容易出bug
			return err
		}
		if n != 8 {
			return errors.New("micro: 没读够数据")
		}

		res := handleMsg(bs)
		n, err = conn.Write(res)
		if err != nil { // 老师建议错误了就返回, 简单粗暴，不容易出bug
			return err
		}
		if n != len(res) {
			return errors.New("micro: 没写完数据")
		}

	}
}

func handleMsg(req []byte) []byte {
	res := make([]byte, len(req)*2)
	copy(res[:len(req)], req)
	copy(res[len(req):], req)
	return res
}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(network, addr string) error {
	listener, err := net.Listen(network, addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			// 一般是端口被占用的错误
			return err
		}

		go func() {
			if err := s.handleConn(conn); err != nil {
				_ = conn.Close()
			}
		}()
	}
}

// 我们可以认为一个请求包含2个部分
// 1. 长度字段：8个字节
// 2. 请求数据
// 响应也是这个规范
func (s *Server) handleConn(conn net.Conn) error {
	for {
		// lenBs 是长度字段的字节表示
		lenBs := make([]byte, numOfLengthBytes)
		_, err := conn.Read(lenBs)
		if err != nil {
			return err
		}

		// 消息有多长？
		length := binary.BigEndian.Uint64(lenBs)
		reqBs := make([]byte, length)
		n, err := conn.Read(reqBs)
		if err != nil {
			return err
		}
		if uint64(n) != length {
			return fmt.Errorf("micro: 数据读取不完整")
		}
		respData := handleMsg(reqBs)
		respLen := len(respData)
		// 我要在这，构建响应数据
		// data = respLen 的64位表示 + respData
		res := make([]byte, respLen+numOfLengthBytes)

		// 第一步: 把长度写进去前 numOfLengthBytes 个字节
		binary.BigEndian.PutUint64(res[:numOfLengthBytes], uint64(respLen))
		// 第二步: 写入数据
		copy(res[numOfLengthBytes:], respData)
		_, err = conn.Write(res)
		if err != nil {
			return err
		}

	}
}

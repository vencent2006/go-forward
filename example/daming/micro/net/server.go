package net

import (
	"errors"
	"io"
	"net"
)

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
		n, err := conn.Read(bs)
		if err == net.ErrClosed || err == io.EOF || err == io.ErrUnexpectedEOF {
			// 这种是明确关闭的连接
			return err
		}
		if err != nil { // 这种是可以挽救的错误
			continue
		}
		if n != 8 {
			return errors.New("micro: 没读够数据")
		}

		res := handleMsg(bs)
		n, err = conn.Write(res)
		if err == net.ErrClosed || err == io.EOF || err == io.ErrUnexpectedEOF {
			// 这种是明确关闭的连接
			return err
		}
		if err != nil { // 这种是可以挽救的错误
			continue
		}
		if n != len(res) {
			return errors.New("micro: 没写完数据")
		}

	}
}

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

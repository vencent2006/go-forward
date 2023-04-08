package net

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func Connect(network, addr string) error {
	conn, err := net.DialTimeout(network, addr, time.Second*3)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()

	for i := 0; i < 10; i++ {
		_, err = conn.Write([]byte("hello"))
		if err != nil {
			return err
		}

		res := make([]byte, 128)
		n, err := conn.Read(res)
		if err != nil {
			return err
		}
		fmt.Println(string(res[:n]))
	}

	return nil
}

type Client struct {
	network string
	addr    string
}

func NewClient(network, addr string) *Client {
	return &Client{
		network: network,
		addr:    addr,
	}
}

func (c *Client) Send(data string) (string, error) {
	conn, err := net.DialTimeout(c.network, c.addr, time.Second*3)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = conn.Close()
	}()

	reqLen := len(data)
	// 我要在这，构建响应数据
	// data = reqLen 的64位表示 + respData
	req := make([]byte, reqLen+numOfLengthBytes)

	// 第一步: 把长度写进去前 numOfLengthBytes 个字节
	binary.BigEndian.PutUint64(req[:numOfLengthBytes], uint64(reqLen))
	// 第二步: 写入数据
	copy(req[numOfLengthBytes:], data)
	_, err = conn.Write(req)
	if err != nil {
		return "", err
	}

	// lenBs 是长度字段的字节表示
	lenBs := make([]byte, numOfLengthBytes)
	_, err = conn.Read(lenBs)
	if err != nil {
		return "", err
	}

	// 消息有多长？
	length := binary.BigEndian.Uint64(lenBs)
	respBs := make([]byte, length)

	_, err = conn.Read(respBs)
	if err != nil {
		return "", err
	}

	return string(respBs), nil
}

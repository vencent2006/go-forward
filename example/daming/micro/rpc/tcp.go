package rpc

import (
	"encoding/binary"
	"net"
)

const numOfLengthBytes = 8

func ReadMsg(conn net.Conn) ([]byte, error) {
	// 协议头和协议体
	// lenBs 是长度字段的字节表示
	lenBs := make([]byte, numOfLengthBytes)

	_, err := conn.Read(lenBs)
	if err != nil {
		return nil, err
	}
	// 消息有多长？
	headerLength := binary.BigEndian.Uint32(lenBs[:4])
	bodyLength := binary.BigEndian.Uint32(lenBs[4:])
	length := headerLength + bodyLength
	data := make([]byte, length)
	_, err = conn.Read(data[8:])
	copy(data[:8], lenBs)
	return data, err
}

func EncodeMsg(data []byte) []byte {
	respLen := len(data)
	res := make([]byte, respLen+numOfLengthBytes) // data = respLen 的64位表示 + respData
	// 第一步: 把长度写进去前 numOfLengthBytes 个字节
	binary.BigEndian.PutUint64(res[:numOfLengthBytes], uint64(respLen))
	// 第二步: 写入数据
	copy(res[numOfLengthBytes:], data)
	return res
}

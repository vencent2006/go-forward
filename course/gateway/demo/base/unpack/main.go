package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const Msg_Header = "12345678"

func main() {
	// 缓冲区
	bytesBuffer := bytes.NewBuffer([]byte{})

	// 发送
	if err := Encode(bytesBuffer, "hello world 0!!!"); err != nil {
		panic(err)
	}
	if err := Encode(bytesBuffer, "hello world 1!!!"); err != nil {
		panic(err)
	}

	// 接收
	for {
		if bt, err := Decode(bytesBuffer); err == nil {
			// no exception
			fmt.Println(string(bt))
			continue
		}
		break
	}
}

// Encode 编码
func Encode(bytesBuffer io.Writer, content string) error {
	// msg_header + content_len + content
	// 8          + 4           + len(content)

	// msg_header
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(Msg_Header)); err != nil {
		return err
	}

	// content len
	clen := int32(len([]byte(content)))
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}

	// content
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}

	return nil
}

// Decode 解码
func Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	// msg_header + content_len + content
	// 8          + 4           + len(content)

	// msg_header
	MagicBuf := make([]byte, len(Msg_Header))
	if _, err = io.ReadFull(bytesBuffer, MagicBuf); err != nil {
		return nil, err
	}

	// content_len
	lengthBuf := make([]byte, 4) // 对应int32
	if _, err = io.ReadFull(bytesBuffer, lengthBuf); err != nil {
		return nil, err
	}

	// content
	length := binary.BigEndian.Uint32(lengthBuf)
	bodyBuf = make([]byte, length)
	if _, err = io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err
	}

	return bodyBuf, err
}

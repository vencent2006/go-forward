package aaa

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const msgHeader = "12345678"

func Foo() {
	fmt.Println("hello world")
}

func Encode(bytesBuffer io.Writer, content string) error {
	// header + contentLen + content
	// 8 + 4 + contentLen

	// 1. header
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(msgHeader)); err != nil {
		return err
	}

	// 2. contentLen
	contentLen := int32(len(content))
	if err := binary.Write(bytesBuffer, binary.BigEndian, contentLen); err != nil {
		return err
	}

	// 3. content
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}

	return nil
}

func Decode(bytesBuffer io.Reader) ([]byte, error) {
	magicBuf := make([]byte, len(msgHeader))
	var err error

	// 1. msgHeader
	if _, err = io.ReadFull(bytesBuffer, magicBuf); err != nil {
		return nil, err
	}

	// check header
	if string(magicBuf) != msgHeader {
		return nil, errors.New("msgHeader mismatch")
	}

	// 2. contentLen
	lengthBuf := make([]byte, 4)
	if _, err = io.ReadFull(bytesBuffer, lengthBuf); err != nil {
		return nil, err
	}

	contentLen := binary.BigEndian.Uint32(lengthBuf)
	bodyBuf := make([]byte, contentLen)
	if _, err = io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err
	}

	return bodyBuf, nil
}

package parser

import (
	"bufio"
	"errors"
	"go-examples/course/go-redis-rewrite/go-redis/interface/resp"
	"io"
	"strconv"
)

type Payload struct {
	Data resp.Reply
	Err  error
}

type readState struct {
	readingMultiLine  bool     // 单行还是多行
	expectedArgsCount int      // 应该解析的参数个数
	msgType           byte     // 消息类型
	args              [][]byte // 已经解析的参数
	bulkLen           int64    // 字节的长度
}

// 解析已经结束
func (s *readState) finished() bool {
	return s.expectedArgsCount > 0 && len(s.args) == s.expectedArgsCount
}

func ParseStream(reader io.Reader) <-chan *Payload {
	ch := make(chan *Payload)
	go parse0(reader, ch)
	return ch
}

// 解析器的核心
func parse0(reader io.Reader, ch chan<- *Payload) {

}

// *3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n
func readLine(bufReader *bufio.Reader, state *readState) ([]byte, bool, error) {
	//  1. \r\n切分
	//	2. 之前读到了$数字，严格读取字符个数
	var msg []byte
	var err error
	if state.bulkLen == 0 { // 1. \r\n切分
		msg, err = bufReader.ReadBytes('\n')
		if err != nil {
			return nil, true, err
		}
		if len(msg) == 0 || msg[len(msg)-2] != '\r' {
			//长度和倒数第二个字符
			return nil, false, errors.New("protocol error: " + string(msg))
		}
	} else { // 2.之前读到了$数字，严格读取字符个数
		// state.bulkLen > 0
		msg = make([]byte, state.bulkLen+2) // 2代表\r\n
		_, err := io.ReadFull(bufReader, msg)
		if err != nil {
			return nil, true, err
		}
		if len(msg) == 0 || msg[len(msg)-2] != '\r' || msg[len(msg)-1] != '\n' {
			return nil, false, errors.New("protocol error: " + string(msg))
		}
		state.bulkLen = 0
	}
	return msg, false, nil
}

func parseMultiBulkHeader(msg []byte, state *readState) error {
	var err error
	var expectedLine uint64
	expectedLine, err = strconv.ParseUint(string(msg[1:len(msg)-2]), 10, 32)
	if err != nil {
		return errors.New("protocol error: " + string(msg))
	}
	if expectedLine == 0 {
		state.expectedArgsCount = 0
		return nil
	} else if expectedLine > 0 {
		state.msgType = msg[0]
		state.readingMultiLine = true
		state.expectedArgsCount = int(expectedLine)
		state.args = make([][]byte, 0, expectedLine)
		return nil
	} else {
		return errors.New("protocol error: " + string(msg))
	}
}

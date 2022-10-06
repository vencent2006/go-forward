package parser

import (
	"go-examples/course/go-redis-rewrite/go-redis/interface/resp"
	"io"
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

package reply

var unknownErrBytes = []byte("-Err unknown\r\n")

type UnknownErrReply struct{}

func (u UnknownErrReply) Error() string {
	return "Err unknown"
}

func (u UnknownErrReply) ToBytes() []byte {
	return unknownErrBytes
}

// ArgNumErrReply 参数错误
type ArgNumErrReply struct {
	Cmd string
}

func (r ArgNumErrReply) Error() string {
	panic("implement me")
}

func (r *ArgNumErrReply) ToBytes() []byte {
	return []byte("-ERR wrong number of arguments for '" + r.Cmd + "' command\r\n")
}

func MakeArgNumErrReply(cmd string) *ArgNumErrReply {
	return &ArgNumErrReply{Cmd: cmd}
}

// SyntaxErrReply 语法错误
type SyntaxErrReply struct{}

func (s SyntaxErrReply) Error() string {
	return "Err syntax err"
}

func (s SyntaxErrReply) ToBytes() []byte {
	return syntaxErrBytes
}

var syntaxErrBytes = []byte("-Err syntax err\r\n")
var theSyntaxErrReply = &SyntaxErrReply{}

func MakeSyntaxErrReply() *SyntaxErrReply {
	return theSyntaxErrReply
}

// 类型错误
var wrongTypeErrBytes = []byte("-WRONGTYPE Operation against a key holding the wrong kind of value\r\n")

type WrongTypeErrReply struct{}

func (r *WrongTypeErrReply) ToBytes() []byte {
	return wrongTypeErrBytes
}

func (r *WrongTypeErrReply) Error() string {
	return "Operation against a key holding the wrong kind of value"
}

// ProtocolErrReply 接口协议错误
type ProtocolErrReply struct {
	Msg string
}

func (r *ProtocolErrReply) Error() string {
	return "-ERR Protocol error: '" + r.Msg + "'"
}

func (r *ProtocolErrReply) ToBytes() []byte {
	return []byte("-ERR Protocol error: '" + r.Msg + "'\r\n")
}

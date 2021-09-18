/**
 * @Author: vincent
 * @Description:
 * @File:  opcode
 * @Version: 1.0.0
 * @Date: 2021/8/12 16:15
 */

package kim

import "time"

const (
	DefaultReadWait  = time.Minute * 3
	DefaultWriteWait = time.Second * 10
	DefaultLoginWait = time.Second * 10
	DefaultHeartbeat = time.Second * 55
)

type OpCode byte

const (
	OpContinuation OpCode = 0x00
	OpText         OpCode = 0x01
	OpBinary       OpCode = 0x02
	OpClose        OpCode = 0x08
	OpPing         OpCode = 0x09
	OpPong         OpCode = 0x0a
)

type Frame interface {
	// opCode
	SetOpCode(code OpCode)
	GetOpCode() OpCode

	// payload
	SetPayload([]byte)
	GetPayload() []byte
}

/**
 * @Author: vincent
 * @Description:
 * @File:  log
 * @Version: 1.0.0
 * @Date: 2021/11/2 17:07
 */

package contract

import (
	"context"
	"io"
	"time"
)

const LogKey = "hade:log"

type LogLevel uint32

const (
	// UnknownLevel 表示未知的日记级别
	UnknownLevel LogLevel = iota
	PanicLevel            // panic 表示会导致整个程序出现崩溃的日志信息
	FatalLevel            // fatal 表示会导致当前这个请求出现提前终止的错误信息
	ErrorLevel            // error 表示出现错误，但是不一定影响后续请求逻辑的错误信息
	WarnLevel             // warn  表示出现错误，但是一定不影响后续请求逻辑的报警信息
	InfoLevel             // info  表示正常的日志信息输出
	DebugLevel            // debug 表示在调试状态下打印出来的日志信息
	TraceLevel            // trace 表示最详细的信息，一般信息量比较大，可能包含调用堆栈等信息
)

// CtxFielder 定义了从context中获取信息的方法
type CtxFielder func(ctx context.Context) map[string]interface{}

// Formatter 定义了将日志信息组织成字符串的通用方法
type Formatter func(level LogLevel, t time.Time, msg string, fields map[string]interface{}) ([]byte, error)

// Log 定义了日志服务协议
type Log interface {
	Panic(ctx context.Context, msg string, fields map[string]interface{})
	Fatal(ctx context.Context, msg string, fields map[string]interface{})
	Error(ctx context.Context, msg string, fields map[string]interface{})
	Warn(ctx context.Context, msg string, fields map[string]interface{})
	Info(ctx context.Context, msg string, fields map[string]interface{})
	Debug(ctx context.Context, msg string, fields map[string]interface{})
	Trace(ctx context.Context, msg string, fields map[string]interface{})

	// 设置日志级别
	SetLevel(level LogLevel)
	// 从context中获取上下文字段field
	SetCtxFielder(handler CtxFielder)
	// 设置输出格式
	SetFormatter(formatter Formatter)
	// 设置输出管道
	SetOutput(out io.Writer)
}

/**
 * @Author: vincent
 * @Description:
 * @File:  context
 * @Version: 1.0.0
 * @Date: 2021/9/21 22:24
 */

package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context

	// 是否超时标记位
	hasTimeout bool
	writerMux  *sync.Mutex // 写保护机制

	// 当前请求的handler链条
	handlers []ControllerHandler
	index    int // 当前请求调用到调用链的哪个节点

	// url路由匹配的参数
	params map[string]string
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(), // 包含在request中
		writerMux:      &sync.Mutex{},
		index:          -1, // 这个初值非常重要
	}
}

// ------ region begin: base function
func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// 为context设置handlers
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

// 设置参数
func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

// 核心函数，调用context的下一个函
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// ------ region end: base function

// ------ region begin: implement context.Context
func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

// ------ region end: implement context.Context

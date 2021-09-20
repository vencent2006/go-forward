/**
 * @Author: vincent
 * @Description:
 * @File:  context
 * @Version: 1.0.0
 * @Date: 2021/9/20 01:11
 */

package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handler        ControllerHandler

	// 是否超时标记位
	hasTimeout bool
	// 写保护机制
	writerMux *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(), // 包含在request中
		writerMux:      &sync.Mutex{},
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

func (ctx *Context) SetHandler(handler ControllerHandler) {
	ctx.handler = handler
}

func (ctx *Context) SetHasTimout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
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

// ------ region begin: request
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.URL.Query())
	}
	return map[string][]string{}
}

// query

func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		// existed
		length := len(values)
		if length > 0 {
			// 取最后一个
			intVal, err := strconv.Atoi(values[length-1])
			if err != nil {
				// 返回default
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		length := len(values)
		if length > 0 {
			return values[length-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}

// post
func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.PostForm)
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		length := len(values)
		if length > 0 {
			intVal, err := strconv.Atoi(values[length-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		length := len(values)
		if length > 0 {
			return values[length-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}

// application/json post
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}

		// todo: ioutil.NopCloser是个啥
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// ------ region end: request

// ------ region begin: response
func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}

	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)

	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}

	ctx.responseWriter.Write(byt)
	return nil
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}

// ------ region end: response

/**
 * @Author: vincent
 * @Description:
 * @File:  hade_context
 * @Version: 1.0.0
 * @Date: 2021/10/27 17:20
 */

package gin

import (
	"context"
	"go-examples/course/handwriting-web-inf/code_12/framework"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

// 绑定是全局的，所以在Engine中Bind和IsBind
// engine实现container的绑定封装
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

// IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

// ----- context实现container的几个封装
// 实例化是在请求中的，所以在Context中Make
// 实现make的封装
func (ctx *Context) Make(key string) (interface{}, error) {
	return ctx.container.Make(key)
}

// 实现mustMake的封装
func (ctx *Context) MustMake(key string) interface{} {
	return ctx.container.MustMake(key)
}

// 实现makeNew的封装
func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ctx.container.MakeNew(key, params)
}

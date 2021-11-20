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
)

func (c *Context) BaseContext() context.Context {
	return c.Request.Context()
}

// ----- context实现container的几个封装
// 实例化是在请求中的，所以在Context中Make
// 实现make的封装
func (c *Context) Make(key string) (interface{}, error) {
	return c.container.Make(key)
}

// 实现mustMake的封装
func (c *Context) MustMake(key string) interface{} {
	return c.container.MustMake(key)
}

// 实现makeNew的封装
func (c *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return c.container.MakeNew(key, params)
}

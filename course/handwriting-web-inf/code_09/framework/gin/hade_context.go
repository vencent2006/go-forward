/**
 * @Author: vincent
 * @Description:
 * @File:  hade_context
 * @Version: 1.0.0
 * @Date: 2021/10/27 17:20
 */

package gin

import "context"

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

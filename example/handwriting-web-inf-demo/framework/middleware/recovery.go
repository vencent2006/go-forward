/**
 * @Author: vincent
 * @Description:
 * @File:  recovery
 * @Version: 1.0.0
 * @Date: 2021/10/26 09:42
 */

package middleware

import "go-examples/example/handwriting-web-inf-demo/framework"

// recovery机制，将协程中的函数异常进行捕获
func Recovery() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		// 核心在增加这个recover机制，捕获c.Next()出现的panic
		defer func() {
			if err := recover(); err != nil {
				c.Json(500, err)
			}
		}()
		// 使用next执行具体的业务逻辑
		c.Next()

		return nil
	}
}

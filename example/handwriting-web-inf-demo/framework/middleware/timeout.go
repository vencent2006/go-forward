/**
 * @Author: vincent
 * @Description:
 * @File:  timeout
 * @Version: 1.0.0
 * @Date: 2021/10/26 09:48
 */

package middleware

import (
	"context"
	"fmt"
	"go-examples/example/handwriting-web-inf-demo/framework"
	"log"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		// 执行业务前的预操作: 初始化超时context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel() // cancel是通知子context的

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			// 使用next执行具体的业务逻辑
			c.Next()

			finish <- struct{}{}
		}()

		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			// todo: 这个返回timeout很神奇
			c.Json(500, "time out")
			log.Println(p)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			//收到父context的执行完毕通知
			c.SetHasTimeout()
			c.Json(500, "time out 2")
		}

		return nil
	}
}

/**
 * @Author: vincent
 * @Description:
 * @File:  timeout
 * @Version: 1.0.0
 * @Date: 2021/9/21 23:28
 */

package middleware

import (
	"context"
	"fmt"
	"go-examples/course/handwriting-web-inf/code_21/framework/gin"
	"log"
	"time"
)

func Timeout(d time.Duration) gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		// 执行业务逻辑前预操作：初始化超时context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

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
			c.ISetStatus(500).IJson("panic")
			log.Println(p)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.ISetStatus(500).IJson("time out")
		}
	}
}

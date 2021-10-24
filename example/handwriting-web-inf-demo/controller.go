/**
 * @Author: vincent
 * @Description:
 * @File:  controller
 * @Version: 1.0.0
 * @Date: 2021/10/24 16:25
 */

package main

import (
	"context"
	"fmt"
	"go-examples/example/handwriting-web-inf-demo/framework"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	// 1秒后超时
	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// do real action
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		finish <- struct{}{}
	}()

	// 父goroutine
	select {
	case p := <-panicChan:
		c.WriteMux().Lock()
		defer c.WriteMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		// finish，就已经调用ctx.Json了
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriteMux().Lock()
		defer c.WriteMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}

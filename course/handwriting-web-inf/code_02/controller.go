/**
 * @Author: vincent
 * @Description:
 * @File:  controller
 * @Version: 1.0.0
 * @Date: 2021/9/20 14:50
 */

package main

import (
	"context"
	"fmt"
	"go-examples/course/handwriting-web-inf/code_02/framework"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	finishChan := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		// Do real action
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		// push finish
		finishChan <- struct{}{}
	}()

	// select
	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finishChan:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimout()
	}
	return nil
}

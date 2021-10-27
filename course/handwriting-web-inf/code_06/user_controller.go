/**
 * @Author: vincent
 * @Description:
 * @File:  user_controller
 * @Version: 1.0.0
 * @Date: 2021/9/21 23:30
 */

package main

import (
	"go-examples/course/handwriting-web-inf/code_06/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)
	// 输出结果
	c.SetOkStatus().Json("ok, UserLoginController" + foo)
	return nil
}

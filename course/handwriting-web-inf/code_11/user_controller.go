/**
 * @Author: vincent
 * @Description:
 * @File:  user_controller
 * @Version: 1.0.0
 * @Date: 2021/9/21 23:30
 */

package main

import (
	"go-examples/course/handwriting-web-inf/code_11/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context) {
	foo, _ := c.DefaultQueryString("foo", "def")
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)
	// 输出结果
	c.ISetOkStatus().IJson("ok, UserLoginController" + foo)
}

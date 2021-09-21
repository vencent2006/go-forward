/**
 * @Author: vincent
 * @Description:
 * @File:  user_controller
 * @Version: 1.0.0
 * @Date: 2021/9/21 23:30
 */

package main

import "go-examples/course/handwriting-web-inf/code_04/framework"

func UserLoginController(c *framework.Context) error {
	c.Json(200, "ok, UserLoginController")
	return nil
}

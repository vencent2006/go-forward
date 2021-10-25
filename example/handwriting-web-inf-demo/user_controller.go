/**
 * @Author: vincent
 * @Description:
 * @File:  user_controller
 * @Version: 1.0.0
 * @Date: 2021/10/25 20:00
 */

package main

import "go-examples/example/handwriting-web-inf-demo/framework"

func UserLoginController(c *framework.Context) error {
	c.Json(200, "ok, UserLoginController")
	return nil
}

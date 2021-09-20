package main

import "go-examples/course/handwriting-web-inf/code_03/framework"

func UserLoginController(c *framework.Context) error {
	c.Json(200, "ok, UserLoginController")
	return nil
}

/**
 * @Author: vincent
 * @Description:
 * @File:  test
 * @Version: 1.0.0
 * @Date: 2021/10/26 09:45
 */

package middleware

import (
	"fmt"
	"go-examples/example/handwriting-web-inf-demo/framework"
)

func Test1() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
		return nil
	}
}

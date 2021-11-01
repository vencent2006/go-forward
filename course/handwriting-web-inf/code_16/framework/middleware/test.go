/**
 * @Author: vincent
 * @Description:
 * @File:  test
 * @Version: 1.0.0
 * @Date: 2021/9/21 23:28
 */

package middleware

import (
	"fmt"
	"go-examples/course/handwriting-web-inf/code_16/framework/gin"
)

func Test1() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware post test1")
	}
}

func Test2() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware post test2")
	}
}

func Test3() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
	}
}

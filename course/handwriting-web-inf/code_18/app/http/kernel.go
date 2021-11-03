/**
 * @Author: vincent
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2021/10/29 15:38
 */

package http

import (
	"go-examples/course/handwriting-web-inf/code_18/framework"
	"go-examples/course/handwriting-web-inf/code_18/framework/gin"
)

// NewHttpEngine is command
func NewHttpEngine(container framework.Container) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// todo: vincent添加的传递container
	r.SetContainer(container)

	Routes(r)
	return r, nil
}

/**
 * @Author: vincent
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2021/10/29 15:38
 */

package http

import "go-examples/course/handwriting-web-inf/code_12/framework/gin"

// NewHttpEngine is command
func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	Routes(r)
	return r, nil
}

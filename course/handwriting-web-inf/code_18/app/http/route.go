/**
 * @Author: vincent
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2021/10/29 15:39
 */

package http

import (
	"go-examples/course/handwriting-web-inf/code_18/app/http/module/demo"
	"go-examples/course/handwriting-web-inf/code_18/framework/gin"
)

func Routes(r *gin.Engine) {
	demo.Register(r)
}

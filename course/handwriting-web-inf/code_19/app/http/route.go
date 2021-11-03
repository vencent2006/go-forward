/**
 * @Author: vincent
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2021/10/29 15:39
 */

package http

import (
	"go-examples/course/handwriting-web-inf/code_19/app/http/module/demo"
	"go-examples/course/handwriting-web-inf/code_19/framework/gin"
	"go-examples/course/handwriting-web-inf/code_19/framework/middleware/static"
)

func Routes(r *gin.Engine) {
	// 根路径"/"先去"/frontend/dist"目录下，查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	// 动态路由定义
	demo.Register(r)
}

/**
 * @Author: vincent
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2021/10/29 15:39
 */

package http

import (
	"go-examples/course/handwriting-web-inf/code_29/app/http/module/demo"
	"go-examples/course/handwriting-web-inf/code_29/framework/contract"
	"go-examples/course/handwriting-web-inf/code_29/framework/gin"
	ginSwagger "go-examples/course/handwriting-web-inf/code_29/framework/middleware/gin-swagger"
	"go-examples/course/handwriting-web-inf/code_29/framework/middleware/gin-swagger/swaggerFiles"
	"go-examples/course/handwriting-web-inf/code_29/framework/middleware/static"
	"log"
)

func Routes(r *gin.Engine) {
	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	// 根路径"/"先去"/frontend/dist"目录下，查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	// 如果配置了swagger, 则显示swagger的中间件
	log.Println("app.swagger is ", configService.GetBool("app.swagger"))
	if configService.GetBool("app.swagger") == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 动态路由定义
	demo.Register(r)
}

/**
 * @Author: vincent
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2021/10/24 16:24
 */

package main

import (
	"go-examples/example/handwriting-web-inf-demo/framework"
	"go-examples/example/handwriting-web-inf-demo/framework/middleware"
)

func registerRouter(core *framework.Core) {
	// 静态路由+HTTP方法匹配
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test2())
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", middleware.Test2(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}

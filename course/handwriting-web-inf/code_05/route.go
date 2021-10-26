/**
 * @Author: vincent
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2021/9/21 23:31
 */

package main

import (
	"go-examples/course/handwriting-web-inf/code_05/framework"
	"go-examples/course/handwriting-web-inf/code_05/framework/middleware"
)

// 注册路由规则
func registerRouter(core *framework.Core) {
	// 静态路由+HTTP方法匹配
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test3())
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}

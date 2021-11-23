/**
 * @Author: vincent
 * @Description:
 * @File:  env
 * @Version: 1.0.0
 * @Date: 2021/10/29 10:01
 */

package test

import (
	"go-examples/course/handwriting-web-inf/code_29/framework"
	"go-examples/course/handwriting-web-inf/code_29/framework/provider/app"
	"go-examples/course/handwriting-web-inf/code_29/framework/provider/env"
)

const (
	BasePath = "/Users/mingzhe8/Documents/innovation/go_proj_exercise/vincent/go-forward/course/handwriting-web-inf/code_29/"
)

func InitBaseContainer() framework.Container {
	// 初始化服务容器
	container := framework.NewHadeContainer()
	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{BaseFolder: BasePath})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.HadeTestingEnvProvider{})
	return container
}

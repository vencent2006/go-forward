package main

import (
	"go-examples/course/handwriting-web-inf/code_28/app/console"
	"go-examples/course/handwriting-web-inf/code_28/app/http"
	"go-examples/course/handwriting-web-inf/code_28/framework"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/app"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/cache"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/config"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/distributed"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/env"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/kernel"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/log"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/orm"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/redis"
)

func main() {

	// 初始化服务容器
	container := framework.NewHadeContainer()
	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.HadeEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.HadeConfigProvider{})
	container.Bind(&log.HadeLogServiceProvider{})
	container.Bind(&orm.GormProvider{})
	container.Bind(&redis.RedisProvider{})
	container.Bind(&cache.HadeCacheProvider{})

	// 后续初始化需要绑定的服务提供者...

	// 将HTTP引擎初始化，并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}
	//container.Bind(&demoService.DemoProvider{})

	// 运行root命令
	console.RunCommand(container)
}

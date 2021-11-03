/**
 * @Author: vincent
 * @Description:
 * @File:  hade_engine
 * @Version: 1.0.0
 * @Date: 2021/10/29 20:51
 */

package gin

import (
	"fmt"
	"go-examples/course/handwriting-web-inf/code_19/framework"
)

func (engine *Engine) SetContainer(container framework.Container) {
	engine.container = container
}

// 绑定是全局的，所以在Engine中Bind和IsBind
// engine实现container的绑定封装
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	fmt.Printf("gin.Engine Bind() container(%p) provider(%v)\n", engine.container, provider)
	return engine.container.Bind(provider)
}

// IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

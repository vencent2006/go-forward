/**
 * @Author: vincent
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/10/29 15:03
 */

package demo

import "go-examples/course/handwriting-web-inf/code_24/framework"

// 服务提供方
type DemoProvider struct {
	// 这是要实现的interface
	framework.ServiceProvider

	// 这是参数
	c framework.Container
}

// Name方法直接将服务对应的字符串凭证返回，在这个例子中就是DemoKey:"demo"
func (sp *DemoProvider) Name() string {
	return DemoKey
}

// Register方法是注册初始化服务实例的方法，我们这里先暂定为NewDemoService
func (sp *DemoProvider) Register(c framework.Container) framework.NewInstance {
	return NewService
}

// IsDefer方法表示是否延迟实例化,我们这里设置为true，讲这个服务的实例化延迟到第一次make的时候
func (sp *DemoProvider) IsDefer() bool {
	return false
}

// Params方法表示实例化的参数。我们这里只实例化一个参数:container，表示我们在NewDemoService这个函数中，只有一个参数，container
func (sp *DemoProvider) Params(c framework.Container) []interface{} {
	// todo: 这里不用入参的container，也比较奇怪
	return []interface{}{sp.c}
}

// Boot方法我们这里什么逻辑都不执行，只打印一行日志消息
func (sp *DemoProvider) Boot(c framework.Container) error {
	sp.c = c
	return nil
}

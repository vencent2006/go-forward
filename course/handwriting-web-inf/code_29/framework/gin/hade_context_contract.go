/**
 * @Author: vincent
 * @Description:
 * @File:  hade_context_contract
 * @Version: 1.0.0
 * @Date: 2021/11/16 18:04
 */

package gin

import "go-examples/course/handwriting-web-inf/code_29/framework/contract"

// MustMakeApp 从容器中获取App服务
func (c *Context) MustMakeApp() contract.App {
	return c.MustMake(contract.AppKey).(contract.App)
}

// MustMakeKernel 从容器中获取Kernel服务
func (c *Context) MustMakeKernel() contract.Kernel {
	return c.MustMake(contract.KernelKey).(contract.Kernel)
}

// MustMakeConfig 从容器中获取配置服务
func (c *Context) MustMakeConfig() contract.Config {
	return c.MustMake(contract.ConfigKey).(contract.Config)
}

// MustMakeLog 从容器中获取日志服务
func (c *Context) MustMakeLog() contract.Log {
	return c.MustMake(contract.LogKey).(contract.Log)
}

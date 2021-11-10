/**
 * @Author: vincent
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/11/1 10:32
 */

package env

import (
	"go-examples/course/handwriting-web-inf/code_24/framework"
	"go-examples/course/handwriting-web-inf/code_24/framework/contract"
)

type HadeEnvProvider struct {
	Folder string
}

// Register a new function for make a service instance
func (provider *HadeEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeEnv
}

// Boot will be called when the service instantiate
func (provider *HadeEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *HadeEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *HadeEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

// Name define the name for this service
func (provider *HadeEnvProvider) Name() string {
	return contract.EnvKey
}

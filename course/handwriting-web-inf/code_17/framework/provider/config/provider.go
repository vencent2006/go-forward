/**
 * @Author: vincent
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/11/2 09:50
 */

package config

import (
	"go-examples/course/handwriting-web-inf/code_17/framework"
	"go-examples/course/handwriting-web-inf/code_17/framework/contract"
	"path/filepath"
)

type HadeConfigProvider struct {
}

// Register register a new function for make a service instance
func (provider *HadeConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeConfig
}

// Boot will called when the service instantiate
func (provider *HadeConfigProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *HadeConfigProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *HadeConfigProvider) Params(c framework.Container) []interface{} {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()

	// 配置文件夹地址
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)

	return []interface{}{c, envFolder, envService.All()}
}

// Name define the name for this service
func (provider *HadeConfigProvider) Name() string {
	return contract.ConfigKey
}

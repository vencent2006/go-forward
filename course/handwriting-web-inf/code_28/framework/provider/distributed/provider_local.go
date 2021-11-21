/**
 * @Author: vincent
 * @Description:
 * @File:  provider_local
 * @Version: 1.0.0
 * @Date: 2021/10/31 19:07
 */

package distributed

import (
	"go-examples/course/handwriting-web-inf/code_28/framework"
	"go-examples/course/handwriting-web-inf/code_28/framework/contract"
)

// LocalDistributedProvider 提供App的具体实现方法
type LocalDistributedProvider struct {
}

// Register 注册HadeApp方法
func (h *LocalDistributedProvider) Register(container framework.Container) framework.NewInstance {
	return NewLocalDistributedService
}

// Boot 启动调用
func (h *LocalDistributedProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *LocalDistributedProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *LocalDistributedProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

// Name 获取字符串凭证
func (h *LocalDistributedProvider) Name() string {
	return contract.DistributedKey
}

/**
 * @Author: vincent
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/10/29 10:41
 */

package app

import (
	"go-examples/course/handwriting-web-inf/code_27/framework"
	"go-examples/course/handwriting-web-inf/code_27/framework/contract"
)

// HadeAppProvider 提供App的具体实现方法
type HadeAppProvider struct {
	BaseFolder string
}

// Register 注册HadeApp方法
func (h *HadeAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeApp
}

// Boot 启动调用
func (h *HadeAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *HadeAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *HadeAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *HadeAppProvider) Name() string {
	return contract.AppKey
}

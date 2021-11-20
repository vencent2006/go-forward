/**
 * @Author: vincent
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/11/16 16:06
 */

package orm

import (
	"go-examples/course/handwriting-web-inf/code_27/framework"
	"go-examples/course/handwriting-web-inf/code_27/framework/contract"
)

// GormProvider 提供App的具体实现方法
type GormProvider struct {
}

// Register 注册方法
func (p *GormProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeGorm
}

// Boot 启动调用
func (p *GormProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (p *GormProvider) IsDefer() bool {
	return true
}

// Param 获取初始化参数
func (p *GormProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

// Name获取字符串凭证
func (p *GormProvider) Name() string {
	return contract.ORMKey
}

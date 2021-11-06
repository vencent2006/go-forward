/**
 * @Author: vincent
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/10/28 16:30
 */

package demo

import (
	"go-examples/course/handwriting-web-inf/code_21/framework"
)

// 具体的接口实例
type Service struct {
	// 实现接口
	IService

	// 参数
	container framework.Container
}

// 初始化实例的方法
func NewService(params ...interface{}) (interface{}, error) {
	// 这里需要将参数展开
	container := params[0].(framework.Container)

	// 返回实例
	return &Service{container: container}, nil
}

// 实现接口
func (s *Service) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "foo",
		},
		{
			ID:   2,
			Name: "bar",
		},
	}
}

/**
 * @Author: vincent
 * @Description:
 * @File:  contract
 * @Version: 1.0.0
 * @Date: 2021/10/28 16:30
 */

package demo

// Demo服务的key

const Key = "hade:demo"

// Demo服务的接口
type Service interface {
	GetFoo() Foo
}

// Demo服务接口定义的一个数据结构
type Foo struct {
	Name string
}

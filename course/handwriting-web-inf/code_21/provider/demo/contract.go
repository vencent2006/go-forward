/**
 * @Author: vincent
 * @Description:
 * @File:  contract
 * @Version: 1.0.0
 * @Date: 2021/10/28 16:30
 */

package demo

const DemoKey = "demo"

type IService interface {
	GetAllStudent() []Student
}

type Student struct {
	ID   int
	Name string
}

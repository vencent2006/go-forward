/**
 * @Author: vincent
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/8/13 15:17
 */

package naming

import "fmt"

// Service define a Service
type ServiceRegistration interface {
	ServiceID() string
	ServiceName() string

	// ip or domain
	PublicAddress() string
	PublicPort() int
	DialURL() string
	GetProtocol() string
	GetNamespace() string
	GetTags() []string
	GetMeta() map[string]string

	// %v的实现方法
	String() string
}

// DefaultServiceImpl
type DefaultServiceImpl struct {
	Id        string
	Name      string
	Address   string
	Port      int
	Protocol  string
	Namespace string
	Tags      []string
	Meta      map[string]string
}

func (e *DefaultServiceImpl) ServiceID() string {
	return e.Id
}

func (e *DefaultServiceImpl) ServiceName() string {
	return e.Name
}

func (e *DefaultServiceImpl) PublicAddress() string {
	return e.Address
}

func (e *DefaultServiceImpl) PublicPort() int {
	return e.Port
}

func (e *DefaultServiceImpl) DialURL() string {
	if e.Protocol == "tcp" {
		return fmt.Sprintf("%s:%d", e.Address, e.Port)
	}
	return fmt.Sprintf("%s://%s:%d", e.Protocol, e.Address, e.Port)
}

func (e *DefaultServiceImpl) GetProtocol() string {
	return e.Protocol
}

func (e *DefaultServiceImpl) GetNamespace() string {
	return e.Namespace
}

func (e *DefaultServiceImpl) GetTags() []string {
	return e.Tags
}

func (e *DefaultServiceImpl) GetMeta() map[string]string {
	return e.Meta
}

func (e *DefaultServiceImpl) String() string {
	return fmt.Sprintf("Id:%s, Name:%s,Address:%s,Port:%d,Namespace:%s,Tags:%v,Meta:%v",
		e.Id,
		e.Name,
		e.Address,
		e.Port,
		e.Namespace,
		e.Tags,
		e.Meta)
}

func NewEntry(id, name, protocol, address string, port int) ServiceRegistration {
	return &DefaultServiceImpl{
		Id:       id,
		Name:     name,
		Address:  address,
		Port:     port,
		Protocol: protocol,
	}

}

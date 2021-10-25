/**
 * @Author: vincent
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2021/10/25 16:22
 */

package framework

// IGroup 代表前缀分组
type IGroup interface {
	// 实现HttpMethod方法
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	// todo: 实现嵌套group，没有具体的todo，就是关注下这种嵌套的使用方式
	Group(string) IGroup
}

// Group struct 实现了IGroup
type Group struct {
	core   *Core  // 指向core结构
	parent *Group // 指向上一个Group，如果有的话
	prefix string // 这个group的通用前缀
}

func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Get(uri, handler)
}

func (g *Group) Post(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Post(uri, handler)
}

func (g *Group) Put(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Put(uri, handler)
}

func (g *Group) Delete(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Delete(uri, handler)
}

func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{core: core, parent: nil, prefix: prefix}
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		// 找到根儿了
		return g.prefix
	}

	return g.parent.getAbsolutePrefix() + g.prefix
}

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
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)

	// 嵌套结构
	// 嵌套的Group
	Group(string) IGroup
	// 嵌套的中间件
	Use(middlewares ...ControllerHandler)
}

// Group struct 实现了IGroup
type Group struct {
	core   *Core  // 指向core结构
	parent *Group // 指向上一个Group，如果有的话
	prefix string // 这个group的通用前缀

	middlewares []ControllerHandler // 存放中间件
}

func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Get(uri, allHandlers...)
}

func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Post(uri, allHandlers...)
}

func (g *Group) Put(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Put(uri, allHandlers...)
}

func (g *Group) Delete(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Delete(uri, allHandlers...)
}

func (g *Group) Use(middlewares ...ControllerHandler) {
	// 从后面添加middlewares，那么就与添加顺序一致了
	g.middlewares = append(g.middlewares, middlewares...)
}

// 返回IGroup接口
func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:        core,
		parent:      nil,
		prefix:      prefix,
		middlewares: []ControllerHandler{},
	}
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		// 找到根儿了
		return g.prefix
	}

	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		// 父group为空
		return g.middlewares
	}
	// 递归获取父group的middlewares添加到自己的middlewares中
	return append(g.parent.getMiddlewares(), g.middlewares...)
}

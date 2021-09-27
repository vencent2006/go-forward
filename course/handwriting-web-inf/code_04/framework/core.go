/**
 * @Author: vincent
 * @Description:
 * @File:  core
 * @Version: 1.0.0
 * @Date: 2021/9/21 23:06
 */

package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core represent core struct

type Core struct {
	router      map[string]*Tree    // all routers
	middlewares []ControllerHandler // 从core这边设置的中间件
}

func NewCore() *Core {
	// 初始化路由
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

// 注册中间件
func (c *Core) Use(middlewares ...ControllerHandler) {
	// todo: vincent say ，这个是覆盖而不是追加，感觉应该是有问题
	c.middlewares = middlewares
}

func (c *Core) PrintMiddlewares() {
	log.Printf("middlewares is %+v", c.middlewares)
	for i, middleware := range c.middlewares {
		log.Printf("%d: %+v", i, middleware)
	}
}

// === http method wrap begin

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	// 将core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	log.Printf("get %s allHandlers: %+v", url, allHandlers)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// ==== http method wrap end

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler {
	// uri 和 method 全部转化为大写, 保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.ServeHTTP")

	// 封装自定义context
	ctx := NewContext(request, response)

	// 寻找路由
	handlers := c.FindRouteByRequest(request)
	if handlers == nil {
		// 如果没有找到，这里打印日志
		ctx.Json(404, "not found")
		return
	}

	ctx.SetHandlers(handlers)

	// 调用路由函数, 如果返回err，代表存在内部错误，返回500
	if err := ctx.Next(); err != nil {
		ctx.Json(500, "internal error")
		return
	}
}

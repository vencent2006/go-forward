/**
 * @Author: vincent
 * @Description:
 * @File:  core
 * @Version: 1.0.0
 * @Date: 2021/10/23 16:28
 */

package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	// map是要初始化的
	return &Core{router: map[string]ControllerHandler{}}
}

// handler注册路由
func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// todo: implement Handler interface: ServeHttp
	log.Println("core.ServeHTTP")
	ctx := NewContext(request, response)

	// 一个简单的路由选择器,这里直接写死为foo
	router := c.router["foo"]
	if router == nil {
		log.Println("foo router is nil")
		return
	}
	log.Println("core.router")
	// 都是使用foo的controllerHandler来处理的请求
	_ = router(ctx)
}

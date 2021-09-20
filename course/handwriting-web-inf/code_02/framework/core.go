/**
 * @Author: vincent
 * @Description:
 * @File:  core
 * @Version: 1.0.0 * @Date: 2021/9/20 14:38
 */

package framework

import (
	"log"
	"net/http"
)

// Core represent core struct

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) Post(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	return c.router["foo"]
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.ServeHTTP")
	ctx := NewContext(request, response)

	router := c.FindRouteByRequest(request)
	if router == nil {
		return
	}

	log.Println("core.Router")

	ctx.SetHandler(router)

	router(ctx)
}

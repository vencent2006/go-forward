/**
 * @Author: vincent
 * @Description:
 * @File:  core
 * @Version: 1.0.0
 * @Date: 2021/10/23 16:28
 */

package framework

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Core struct {
}

func NewCore() *Core {
	return &Core{}
}
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// todo: implement Handler interface: ServeHttp
	log.Println("come to ServeHttp")

	// 控制器

	obj := map[string]interface{}{
		"data": nil,
	}
	// 设置控制器 response 的 header 部分
	response.Header().Set("Content-Type", "application/json")

	// 从请求体中获取参数
	foo := request.PostFormValue("foo")
	if foo == "" {
		foo = "10"
	}
	fooInt, err := strconv.Atoi(foo)
	if err != nil {
		response.WriteHeader(500)
		return
	}
	// 构建返回结构
	obj["data"] = fooInt
	byt, err := json.Marshal(obj)
	if err != nil {
		response.WriteHeader(500)
		return
	}
	// 构建返回状态，输出返回结构
	response.WriteHeader(200)
	response.Write(byt)

	obj["name"] = "vincent"
	byt, err = json.Marshal(obj)
	if err != nil {
		response.WriteHeader(500)
		return
	}
	// 构建返回状态，输出返回结构
	response.WriteHeader(200)
	response.Write(byt)

	return
}

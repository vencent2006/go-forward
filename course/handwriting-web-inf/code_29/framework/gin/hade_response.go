/**
 * @Author: vincent
 * @Description:
 * @File:  hade_response
 * @Version: 1.0.0
 * @Date: 2021/10/27 17:05
 */

package gin

import (
	"encoding/xml"
	"fmt"
	"go-examples/course/handwriting-web-inf/code_29/framework/gin/internal/json"
	"html/template"
	"net/http"
	"net/url"
)

// IResponse代表返回方法

type IResponse interface {
	// Json输出
	IJson(obj interface{}) IResponse

	// Jsonp输出
	IJsonp(obj interface{}) IResponse

	// xml输出
	IXml(obj interface{}) IResponse

	// html输出
	IHtml(template string, obj interface{}) IResponse

	// string
	IText(format string, values ...interface{}) IResponse

	// 重定向
	IRedirect(path string) IResponse

	// header
	ISetHeader(key string, val string) IResponse

	// Cookie
	ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	ISetStatus(code int) IResponse

	// 设置200状态
	ISetOkStatus() IResponse
}

// Jsonp输出
func (c *Context) IJsonp(obj interface{}) IResponse {
	// 获取请求参数callback
	callbackFunc := c.Query("callback")
	c.ISetHeader("Content-Type", "application/javascript")
	//todo: 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成xss攻击
	callback := template.JSEscapeString(callbackFunc)

	// 输出函数名
	_, err := c.Writer.Write([]byte(callback))
	if err != nil {
		return c
	}
	// 输出左括号
	_, err = c.Writer.Write([]byte("("))
	if err != nil {
		return c
	}

	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return c
	}

	_, err = c.Writer.Write(ret)
	if err != nil {
		return c
	}
	// 输出右括号
	_, err = c.Writer.Write([]byte(")"))
	if err != nil {
		return c
	}
	return c
}

// xml输出
func (c *Context) IXml(obj interface{}) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return c.ISetStatus(http.StatusInternalServerError)
	}
	c.ISetHeader("Content-Type", "application/xml")
	c.Writer.Write(byt)
	return c
}

// html输出
func (c *Context) IHtml(file string, obj interface{}) IResponse {
	// 读取模板文件，创建template实例
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return c
	}

	// 执行Execute方法将obj和模板进行结合
	if err := t.Execute(c.Writer, obj); err != nil {
		return c
	}

	c.ISetHeader("Content-Type", "application/html")
	return c
}

// string
func (c *Context) IText(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	c.ISetHeader("Content-Type", "application/text")
	c.Writer.Write([]byte(out))
	return c
}

// 重定向
func (c *Context) IRedirect(path string) IResponse {
	http.Redirect(c.Writer, c.Request, path, http.StatusMovedPermanently)
	return c
}

// header
func (c *Context) ISetHeader(key string, val string) IResponse {
	c.Writer.Header().Add(key, val)
	return c
}

// Cookie
func (c *Context) ISetCookie(key string, val string, maxAge int, path string, domain string,
	secure bool, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return c
}

// 设置状态码
func (c *Context) ISetStatus(code int) IResponse {
	c.Writer.WriteHeader(code)
	return c
}

// 设置200状态
func (c *Context) ISetOkStatus() IResponse {
	c.Writer.WriteHeader(http.StatusOK)
	return c
}

func (c *Context) IJson(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return c.ISetStatus(http.StatusInternalServerError)
	}
	c.ISetHeader("Content-Type", "application/json")
	c.Writer.Write(byt)
	return c
}

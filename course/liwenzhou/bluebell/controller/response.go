package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code": 10001, // 程序中的错误码
	"msg" : xx,    // 提示信息
	"data": {},    // 数据
}
*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, packResponseInCtx(
		c,
		code,
		code.Msg(),
		nil))
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, packResponseInCtx(
		c,
		code,
		msg,
		nil))
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, packResponseInCtx(
		c,
		CodeSuccess,
		CodeSuccess.Msg(),
		data))
}

func packResponseInCtx(c *gin.Context, code ResCode, msg, data interface{}) *ResponseData {
	r := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	// todo 可能会造成性能问题
	c.Set("responseData", r)
	return r
}

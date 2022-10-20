package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CtxUserIDKey 将当前请求的userID信息保存到请求的上下文c上
const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前登录的用户ID
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getPageInfo(c *gin.Context) (page int64, size int64) {
	// 获取分页参数
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var err error

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}

	// todo page和size的最大值要限制下

	return
}

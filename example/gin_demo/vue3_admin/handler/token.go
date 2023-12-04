package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TokenInfo struct {
	Token string `json:"token"`
}

func HandleLogin(c *gin.Context) {
	r := &Response{
		Code:    CODE_SUCC,
		Data:    TokenInfo{Token: "dc7f7964-60e5-4f61-b65e-0ac43dce6b1f"},
		Message: "登录成功",
		Success: BOOL_SUCC,
	}
	c.JSON(http.StatusOK, r)
}

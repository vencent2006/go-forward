package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

type TokenInfo struct {
	Token string `json:"token"`
}

func main() {
	r := gin.Default()
	r.POST("/api/sys/login", func(c *gin.Context) {
		r := &Response{
			Code:    200,
			Data:    TokenInfo{Token: "dc7f7964-60e5-4f61-b65e-0ac43dce6b1f"},
			Message: "登录成功",
			Success: true,
		}
		c.JSON(http.StatusOK, r)
	})

	r.Run(":8888")
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"msg"`
	Data    map[string]interface{} `json:"data"`
}

func main() {
	r := gin.Default()

	r.GET("/community/get", func(c *gin.Context) {
		data := &Response{
			Code:    10000,
			Message: "succ",
			Data: map[string]interface{}{
				"community_id": "bobo123",
			},
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8911")
}

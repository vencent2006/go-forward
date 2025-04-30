package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    string `json:"data"`
}

func main() {
	r := gin.Default()

	r.GET("/comment", func(c *gin.Context) {
		id := c.Query("id")
		tmpContent := "太温暖、感人了，看得我热泪盈眶。今年其实有在这样去做了，30多了第一次解锁游乐场、第一次看演唱会，利用周末时间出去旅游。虽然今年工作更加忙碌，压力也更大，但还是很开心，感觉生活开始有了属于自己的色彩。明年会继续努力的"
		content := fmt.Sprintf("%s,%s", id, tmpContent)
		data := &Response{
			Code:    10000,
			Message: "succ",
			Data:    content,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8911")
}

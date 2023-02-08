package main

import (
	"go-examples/example/chat_demo_im/logger"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	response := `{"message":"ok"}`
	r.POST("/post", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		if body != nil {
			logger.Printf("body is %+v", string(body))
		}
		c.String(200, response)
	})

	r.Run(":8080")
}

package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("./webhook.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	r := gin.Default()

	r.POST("/wa/notify/charge", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		if body != nil {
			logger.Printf("body is %+v", string(body))
		}
		c.String(200, "ok")
	})

	r.Run(":26001")
}

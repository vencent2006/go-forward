package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var logger *log.Logger

type NotifyChargeRes struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func init() {
	logFile, err := os.OpenFile("./webhook.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	r := gin.Default()

	r.POST("/wallet/nofify/chain/chage", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		if body != nil {
			logger.Printf("body is %+v", string(body))
		}
		data := &NotifyChargeRes{
			Code:    10000,
			Message: "ok",
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8081")
}

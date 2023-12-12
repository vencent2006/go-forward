package main

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
)

func get_result(c *gin.Context) {

	if rand.Intn(10) < 1 {
		c.JSON(http.StatusInternalServerError, "server inner error")
	} else {
		c.JSON(http.StatusOK, "OK")
	}
}

func query_result(c *gin.Context) {

	if rand.Intn(10) < 2 {
		c.JSON(http.StatusInternalServerError, "server inner error")
	} else {
		c.JSON(http.StatusOK, "OK")
	}
}

func main() {

	r := gin.Default()

	r.Use(ginprom.PromMiddleware(nil))

	r.GET("/get_result", get_result)
	r.GET("/query_result", query_result)

	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))

	r.Run(":7777")

}

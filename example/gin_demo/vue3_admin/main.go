package main

import (
	"example/gin_demo/vue3_admin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/sys/login", handler.HandleLogin)
	r.GET("/api/sys/profile", handler.HandleProfile)

	r.Run(":8888")
}

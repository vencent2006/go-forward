package main

import (
	"example/gin_demo/vue3_admin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/sys/login", handler.HandleLogin)
	r.GET("/api/sys/profile", handler.HandleProfile)
	r.GET("/v1/api/logging/suggestions/getorgkeys", handler.GetTenantList)

	r.Run(":8888")
}

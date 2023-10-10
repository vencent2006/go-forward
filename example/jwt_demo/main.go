package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", loginHandler)

	api := r.Group("/api")
	api.Use(jwtAuthMiddleware())
	api.POST("/order", orderHandler)

	r.Run(":9099")
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(c *gin.Context) {
	var req LoginReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, "internal error")
		return
	}
	if req.Username == "admin" && req.Password == "admin" {
		token, err := GenToken(1001)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, map[string]string{"token": token})
		}
		return
	}
	c.JSON(http.StatusForbidden, "forbidden")
}

type OrderReq struct {
	Product string `json:"product"`
	Count   int    `json:"count"`
}

func orderHandler(c *gin.Context) {
	var req OrderReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, "invalid request")
		return
	}
	userId, _ := c.Get("userId")
	greet := fmt.Sprintf("Hi %v, I will give you %v %v", userId, req.Count, req.Product)
	c.JSON(http.StatusOK, greet)
}

func jwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//标准做法： Authorization: Bearer <token>
		//这里简化了
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusForbidden, "empty token")
			c.Abort()
			return
		}
		claim, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusForbidden, err.Error())
			c.Abort()
			return
		}
		c.Set("userId", claim.UserId)
		c.Next()
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// curl -X POST http://127.0.0.1:8080/login -d "name=vincent&password=123"
	r.POST("login", func(c *gin.Context) {
		fmt.Println(c.FullPath())
		username := c.PostForm("name")
		password := c.PostForm("password")
		log.Printf("c.Request.PostForm.Encode(): %s", c.Request.PostForm.Encode())
		c.String(http.StatusOK, username+","+password+" "+c.Request.PostForm.Encode())
	})

	// curl -d '{"name": "emma", "password": "123"}' -H 'Content-Type: application/json' http://127.0.0.1:8080/login2
	r.POST("login2", func(c *gin.Context) {
		fmt.Println(c.FullPath())
		l := new(Login)
		c.BindJSON(&l)
		log.Printf("c.Request.PostForm.Encode(): %s", c.Request.PostForm.Encode())
		c.JSON(http.StatusOK, gin.H{
			"name":     l.Name,
			"password": l.Password,
		})
	})
	r.Run()
}

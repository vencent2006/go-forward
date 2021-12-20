package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("password123456"))
	r.Use(sessions.Sessions("my_session", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Delete("shirdon")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})

	r.Run(":8000")
}

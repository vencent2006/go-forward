package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func func1(c *gin.Context) {
	fmt.Println("func1")
}

func func2(c *gin.Context) {
	fmt.Println("func2 before")
	c.Next()
	fmt.Println("func2 after")
}

func func3(c *gin.Context) {
	fmt.Println("func3")
	//c.Abort()
}

func func4(c *gin.Context) {
	fmt.Println("func4")
	c.Set("name", "q1mi")
}

func func5(c *gin.Context) {
	fmt.Println("func5")
	v, ok := c.Get("name")
	if ok {
		vStr := v.(string) // 类型转换
		fmt.Println(vStr)
	}
}


func main() {
	r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusNotFound,"bad")
	//},func(c *gin.Context) {
	//	c.String(http.StatusOK, "ok")
	//}, func(context *gin.Context) {
	//	fmt.Println("hello world")
	//})
	r.GET("/hello", func(context *gin.Context) {

	})

	shopGroup := r.Group("/shop", func1, func2)
	shopGroup.Use(func3)
	{
		shopGroup.GET("/index", func4, func5)
	}
	r.Run()
}

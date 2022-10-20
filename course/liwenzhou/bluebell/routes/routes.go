package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	// middleware
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler) // 注册
	v1.POST("/login", controller.LoginHandler)   // 登录
	v1.Use(middlewares.JWTAuthMiddleware())      // 中间件
	{
		// community 社区
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:cid", controller.CommunityDetailHandler)
		// post 帖子
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:pid", controller.PostDetailHandler)
		v1.GET("/post", controller.PostListHandler)
	}

	//r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong")
	//})
	// 感觉没用啥用这个，直接报404不就得了
	//r.NoRoute(func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "404",
	//	})
	//})
	return r
}

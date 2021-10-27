package main

import (
	"context"
	"go-examples/course/handwriting-web-inf/code_09/framework/gin"
	"go-examples/course/handwriting-web-inf/code_09/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func main() {
	// logger
	InitLogger()
	defer logger.Sync()

	// core
	core := gin.New()
	core.Use(gin.Recovery())
	core.Use(middleware.Cost())
	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	// 这个goroutine是启动服务的goroutine
	go func() {
		// todo: 如果写成log.Fatal(server.ListenAndServe()), 优雅关闭的机制就不起作用了;还没明白原因
		server.ListenAndServe()
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT(ctrl+c), SIGQUIT(ctrl+\)，SIGTERM(kill非-9)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	log.Println("receive signal")

	// 调用Server.Shutdown graceful结束, 这里设定是5秒超时
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
}

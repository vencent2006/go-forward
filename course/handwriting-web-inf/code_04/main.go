package main

import (
	"go-examples/course/handwriting-web-inf/code_04/framework"
	"go-examples/course/handwriting-web-inf/code_04/framework/middleware"
	"log"
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func main() {
	InitLogger()
	defer logger.Sync()
	core := framework.NewCore()
	// core.Use(
	// 	middleware.Test1(),
	// 	middleware.Test2())
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	// core.Use(middleware.Timeout(1 * time.Second))
	core.PrintMiddlewares()
	registerRouter(core)
	core.PrintMiddlewares()
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Print(err)
	}
}

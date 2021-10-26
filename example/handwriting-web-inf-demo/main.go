package main

import (
	"go-examples/example/handwriting-web-inf-demo/framework"
	"go-examples/example/handwriting-web-inf-demo/framework/middleware"
	"log"
	"net/http"
)

func main() {
	core := framework.NewCore()
	// use middlewares
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	// register router
	registerRouter(core)
	// server ListenAndServe
	server := http.Server{
		Addr:    ":8088",
		Handler: core,
	}
	log.Fatal(server.ListenAndServe())
}

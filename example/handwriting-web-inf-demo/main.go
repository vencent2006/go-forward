package main

import (
	"go-examples/example/handwriting-web-inf-demo/framework"
	"log"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := http.Server{
		Addr:    ":8088",
		Handler: core,
	}
	log.Fatal(server.ListenAndServe())
}

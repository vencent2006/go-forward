package main

import (
	"go-examples/example/handwriting-web-inf-demo/framework"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    ":8088",
		Handler: framework.NewCore(),
	}
	log.Fatal(server.ListenAndServe())
}

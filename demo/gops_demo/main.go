package main

import (
	"log"
	"net/http"

	"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatalf("agent.Listen err:%v", err)
	}

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello, world!"))
	})

	_ = http.ListenAndServe(":6060", http.DefaultServeMux)

}

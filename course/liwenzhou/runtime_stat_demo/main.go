package main

import (
	"log"
	"net/http"

	"github.com/arl/statsviz"
)

func main() {
	statsviz.RegisterDefault()
	log.Println(http.ListenAndServe(":6060", nil))
}
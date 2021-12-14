package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go accessLogTask()
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

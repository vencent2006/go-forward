package main

import (
	"log"
	"net/http"
	"time"
)

const addr = ":1210"

func sayBye(res http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Second * 1)
	res.Write([]byte("bye, bye, this is http server"))
}

func main() {
	// 创建路由器
	mux := http.NewServeMux()
	// 设置路由规则
	mux.HandleFunc("/bye", sayBye)
	server := http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	log.Println("starting http server at " + addr)
	log.Fatal(server.ListenAndServe())
}

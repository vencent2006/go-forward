package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const addr = "127.0.0.1:2002"

func main() {
	rs1 := "http://127.0.0.1:2003/base"
	url1, err := url.Parse(rs1)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url1)
	log.Println("start reverse proxy server at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}

package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
)

const (
	proxy_addr = "http://127.0.0.1:2003"
	port       = "2002"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("start reverse proxy on port: " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	//step 1 解析代理地址，并更改请求体的协议和主机
	proxy, err := url.Parse(proxy_addr)
	if err != nil {
		panic(err)
	}
	r.URL.Scheme = proxy.Scheme
	r.URL.Host = proxy.Host

	//step 2 请求下游
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(r)
	if err != nil {
		log.Println(err)
		return
	}

	//step 3 把下游请求内容返回给上游
	for key, value := range r.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(w)
}

package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
)

var (
	proxy_addr = "http://127.0.0.1:2003"
	port       = "2002"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// step 1 解析代理地址，并更改请求体的协议和主机
	proxy, err := url.Parse(proxy_addr)
	r.URL.Scheme = proxy.Scheme
	r.URL.Host = proxy.Host

	// step 2 请求下游
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(r)
	if err != nil {
		log.Print(err)
		return
	}

	// step 3 把下游请求内容返回上游
	for k, vv := range resp.Header {
		for _, v := range vv {
			// todo: 这个很奇怪呢
			w.Header().Add(k, v)
		}
	}
	defer resp.Body.Close()               // 要把body close
	bufio.NewReader(resp.Body).WriteTo(w) // io.Write(w, resp.Body)
}

func main() {
	// 4-5 反向代理原来如此
	http.HandleFunc("/", handler)
	log.Println("Start serving on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

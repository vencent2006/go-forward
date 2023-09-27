package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

const (
	HEADER_X_FORWARDED_FOR     = "X-Forwarded-For"
	HEADER_X_FORWARDED_FOR_SEP = ","
)

type Pxy struct{}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("receive request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	transport := http.DefaultTransport

	// step 1，浅拷贝对象，然后就再新增属性数据
	outReq := new(http.Request)
	*outReq = *req // 浅拷贝
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err != nil {
		if prior, ok := outReq.Header[HEADER_X_FORWARDED_FOR]; ok {
			// todo 为啥不直接prior里append呢
			clientIP = strings.Join(prior, HEADER_X_FORWARDED_FOR_SEP) + HEADER_X_FORWARDED_FOR_SEP + clientIP
		}
		outReq.Header.Set(HEADER_X_FORWARDED_FOR, clientIP)
	}

	// step 2, 请求下游
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	// step 3, 把下游请求内容返回给上游
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}

	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)
	res.Body.Close()
}

func main() {
	addr := "0.0.0.0:8080"
	fmt.Println("serve on " + addr)
	http.Handle("/", &Pxy{})
	http.ListenAndServe(addr, nil)
}

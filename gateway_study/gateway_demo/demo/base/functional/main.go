package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

// 测试函数是一等公民
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func (f HandlerFunc) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	f(res, req)
}

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello world"))
}

func main() {
	hf := HandlerFunc(HelloHandler)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte("test")))

	hf.ServeHTTP(res, req)

	bts, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bts))
}

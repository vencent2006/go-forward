package main

import "testing"

const (
	//uri   = "https://fuwu.biz.weibo.com/?1111=2222"
	uri   = "http://127.0.0.1:9002"
	times = 50
)

func Test_long_conn(t *testing.T) {
	longConn(uri, times)
}

func Test_short_conn(t *testing.T) {
	shortConn(uri, times)
}

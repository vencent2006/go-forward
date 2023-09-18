package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// http://www.guoxiaolong.cn/blog/?id=11909

var HTTPTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second, // 连接超时时间
		KeepAlive: 30 * time.Second, // 保持长连接的时间
	}).DialContext, // 设置连接的参数
	MaxIdleConns:          1,                // 最大空闲连接
	IdleConnTimeout:       30 * time.Second, // 空闲连接的超时时间
	ExpectContinueTimeout: 30 * time.Second, // 等待服务第一个响应的超时时间
	MaxIdleConnsPerHost:   100,              // 每个host保持的空闲连接数
}

func main() {
	times := 50
	uri := "https://fuwu.biz.weibo.com/"

	// 短连接的情况
	//shortConn(uri, times)
	// 长连接的情况
	longConn(uri, times)
}

func shortConn(uri string, times int) {
	// 短连接的情况
	start := time.Now()

	for i := 0; i < times; i++ {
		req, err := http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			panic("Http Req Failed " + err.Error())
		}
		client := http.Client{}     // 初始化http的client
		resp, err := client.Do(req) // 发起请求
		if err != nil {
			panic("Http Request Failed " + err.Error())
		}

		time.Sleep(time.Second)
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic("ioutil.ReadAll " + err.Error())
		}
		fmt.Println(i, string(b))
	}
	fmt.Println("Orig GoNet Short Link", time.Since(start))
}

func longConn(uri string, times int) {
	// 长连接的情况
	start2 := time.Now()
	fmt.Printf("transport: %+v\n", HTTPTransport)
	client2 := http.Client{Transport: HTTPTransport} // 初始化一个带有transport的http的client
	for i := 0; i < times; i++ {
		req, err := http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			panic("Http Req Failed " + err.Error())
		}
		resp, err := client2.Do(req)
		if err != nil {
			panic("Http Request Failed " + err.Error())
		}
		defer resp.Body.Close()
		fmt.Println(i, resp.StatusCode)
		ioutil.ReadAll(resp.Body) // 如果不及时从请求中获取结果，此连接会占用，其他请求服务复用连接
		time.Sleep(1 * time.Second)
		fmt.Println(i, "elapse", time.Since(start2))
	}
	fmt.Println("Orig GoNet Long Link", time.Since(start2))
}

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	// 创建连接池
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// 创建客户端
	client := &http.Client{
		Timeout: time.Second*30, // 请求超时时间
		Transport: transport,
	}

	// 请求数据
	resp, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取内容
	bds, err:= ioutil.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}

	fmt.Println(string(bds))

	// 读取内容
}

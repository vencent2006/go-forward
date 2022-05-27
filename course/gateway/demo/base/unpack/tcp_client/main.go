package main

import (
	"fmt"
	"go-examples/course/gateway/demo/base/unpack/unpack"
	"net"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed: %v\n", err)
		return
	}
	for i:=0;i<5;i++{
		content := "hello world " + strconv.Itoa(i)
		unpack.Encode(conn, content)
	}
}

package main

import (
	"fmt"
	"gateway_study/gateway_demo/demo/base/unpack/codec"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed | err:%v", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		bytes, err := codec.Decode(conn)
		if err != nil {
			return
		}
		fmt.Print(string(bytes))
	}
}

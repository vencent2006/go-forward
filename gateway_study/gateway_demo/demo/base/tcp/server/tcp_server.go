package main

import (
	"fmt"
	"net"
)

const addr = "0.0.0.0:9090"

func main() {
	//	1. 监听端口
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed | err:%v\n", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close() //思考题：这里不填写会有啥问题？
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed | err:%v\n", err)
			break
		}
		str := string(buf[:n])
		fmt.Printf("receive from client | data:%s\n", str)
	}
}

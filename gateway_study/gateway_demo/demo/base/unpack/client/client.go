package main

import (
	"bufio"
	"fmt"
	"gateway_study/gateway_demo/demo/base/unpack/codec"
	"net"
	"os"
	"strings"
)

func main() {
	doSend()
	fmt.Print("doSend over")
}

func doSend() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed | err:%v", err)
			return
		}

		if strings.TrimSpace(input) == "Q" {
			return
		}

		err = codec.Encode(conn, input)
		if err != nil {
			fmt.Printf("Encode faile | err:%v", err)
			return
		}
	}
}

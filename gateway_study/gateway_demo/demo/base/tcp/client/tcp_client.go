package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	doSend()
	fmt.Print("doSend over")
}

func doSend() {
	addr := "localhost:9090"
	// 1. 连接服务器
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("conn %s failed | err:%v", addr, err)
		return
	}
	defer conn.Close()

	// 2. 读取命令输入
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 3. 一直读到\n
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from concole failed | err:%v", err)
			break
		}

		// 4. 读到Q停止
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}

		// 5. 回复给服务器
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("write failed | err:%v", err)
			break
		}
	}
}

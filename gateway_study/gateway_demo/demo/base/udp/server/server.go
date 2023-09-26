package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	if err != nil {
		panic(err)
	}

	// for read
	for {
		var data [1024]byte
		n, addr, err := listener.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read from udp failed | err:%v", err)
			break
		}

		go func() {
			fmt.Printf("addr:%v data:%v count:%v\n", addr, string(data[:n]), n)
			_, err := listener.WriteToUDP([]byte("receive success"), addr)
			if err != nil {
				fmt.Printf("WriteToUDP failed | err:%v", err)
			}
		}()
	}
}

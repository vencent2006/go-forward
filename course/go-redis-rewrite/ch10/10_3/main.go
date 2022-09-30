package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("defer main g")

	go func() {
		defer fmt.Println("defer g")
		panic("1111111")
		fmt.Println("end g")
	}()

	time.Sleep(time.Second)
	fmt.Println("end main g")
}

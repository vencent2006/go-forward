package main

import (
	"fmt"
	"time"
)

func do1() {
	do2()
}

func do2() {
	do3()
}

func do3() {
	fmt.Println("do3")
}

func main() {
	go do1()
	time.Sleep(time.Hour)
}

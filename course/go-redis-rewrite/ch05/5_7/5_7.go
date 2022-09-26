package main

import "fmt"

func do1() {
	do2()
}

func do2() {
	do3()
	go do1()

}

func do3() {
	fmt.Println("do3")
}

func main() {
	do1()
	select {}
}

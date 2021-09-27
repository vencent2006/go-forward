package main

import "fmt"

func main() {
	a := 1
	b := 2
	go func() {
		fmt.Println(a, b)
	}()
	a = 99
}

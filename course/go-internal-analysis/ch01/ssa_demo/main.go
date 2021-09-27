package main

var d uint8

func main() {
	var a uint8 = 1
	a = 2
	if true {
		a = 3
	}
	d = a
}

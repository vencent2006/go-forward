package main

func small() string {
	s := "hello, " + "world!"
	return s
}

func fib(index int) int {
	if index < 2 {
		return index
	}
	return fib(index-1) + fib(index-2)
}

func main() {
	small()
	fib(65)
}

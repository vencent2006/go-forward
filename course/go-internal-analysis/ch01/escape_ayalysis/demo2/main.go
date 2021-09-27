package main

func foo() int {
	a := 1
	z := &a
	return *z
}
func main() {
	foo()
}

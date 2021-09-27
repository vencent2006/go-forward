package main

var z *int

func escape() {
	a := 1
	z = &a
}
func main() {
	escape()
	_ = z
}

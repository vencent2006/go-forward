package slice_vs_array

import "testing"

func printSlice(s []int) {
	println(len(s))
	println(cap(s))
}

func printArray(a [100]int) {
	println(len(a))
	println(cap(a))
}

func TestPrint(t *testing.T) {
	s := make([]int, 100)
	printSlice(s)

	var a [100]int
	printArray(a)
}

package main

func printSlice(s []int) {
	print(len(s))
	print(cap(s))
}

func printArray(a [100]int) {
	print(len(a))
	print(cap(a))
}

func main() {
	s := make([]int, 100)
	printSlice(s)

	var a [100]int
	printArray(a)
}

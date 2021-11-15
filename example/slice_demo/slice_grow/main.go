package main

func printSlice(s []int) {
	println(len(s), cap(s))
}

func main() {
	var s []int
	printSlice(s) // 0 0

	s = append(s, 0)
	printSlice(s) // 1 1

	s = append(s, 1)
	s = append(s, 2)
	printSlice(s) // 3 4

	for i := 3; i < 1025; i++ {
		s = append(s, i)
	}
	printSlice(s) // 1025 1280

	var ss []int
	ss = append(ss, 0, 1, 2)
	printSlice(ss)
}

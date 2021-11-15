package main

func main() {
	a := 3
	b := 2
	returnTwo(a, b)
}

func returnTwo(a, b int) (c, d int) {
	tmp := 1 // 这一行的主要目的是保证栈桢不为0，方便分析
	c = a + b
	d = b - tmp
	return
}

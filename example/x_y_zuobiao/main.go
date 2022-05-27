package main

import "fmt"

func main() {
	// 初始值(0,0)
	x := 0
	y := 0
	fmt.Printf("最开始: x = %d, y = %d\n", x, y)

	// x增加10
	x = x + 10
	fmt.Printf("x增加10: x = %d, y = %d\n", x, y)

	// x增加10
	x = x + 10
	fmt.Printf("x增加10: x = %d, y = %d\n", x, y)
	// x设为50
	x = 50
	fmt.Printf("x设为50：x = %d, y = %d\n", x, y)

}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int64
	i = 11
	str := strconv.FormatInt(i, 10)
	if len(str) > 31 {
		number := str[30:31]
		if number == "1" {
			fmt.Println("等于1")
		}
	}
	fmt.Println("here")
}

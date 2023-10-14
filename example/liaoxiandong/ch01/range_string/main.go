package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "programming"
	for i, c := range str {
		fmt.Printf("%2d: %T, %v, %s\n", i, c, c, string(c))
	}
	chars := []rune(str)
	for _, char := range chars {
		fmt.Println(string(char))
	}
	fmt.Println(reflect.TypeOf(str))
}

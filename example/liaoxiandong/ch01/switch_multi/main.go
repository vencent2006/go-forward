package main

import (
	"fmt"
)

func main() {

	var input string
	for {
		fmt.Printf("please input:")
		_, err := fmt.Scan(&input)
		if err != nil {
			panic(err)
		}
		switch input {
		case "java", "python":
			fmt.Println(input, "is a good language")
		case "go", "php":
			fmt.Println(input, "is so so")
		default:
			fmt.Println("i don't know", input)
		}
	}
}

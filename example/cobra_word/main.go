package main

import (
	"log"

	"go-examples/demo/cobra_word/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}

package main

import "fmt"

func main() {
	s := NewServer("1", ":8080")
	defer s.Shutdown()
	err := s.Start()
	fmt.Println(err)
}

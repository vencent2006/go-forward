package main

import "testing"

func TestGrow(t *testing.T) {
	var s []int
	s = append(s, 0, 1, 2)
	println(len(s), cap(s))
}

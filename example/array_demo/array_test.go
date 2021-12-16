package array_demo

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var arr [6]int
	fmt.Printf("----%+v", arr)
	t.Logf("%+v", arr)
}

func TestSlice(t *testing.T) {
	var sliceStr []string
	var sliceNum []int
	var emptySliceNum = []int{}

	fmt.Println(sliceStr == nil)
	fmt.Println(sliceNum == nil)
	fmt.Println(emptySliceNum == nil)
}

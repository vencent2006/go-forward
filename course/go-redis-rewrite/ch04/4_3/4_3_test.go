/**
 * @Author: vincent
 * @Description:
 * @File:  4_2_test
 * @Version: 1.0.0
 * @Date: 2022/9/21 09:44
 */

package main

import (
	"fmt"
	"testing"
)

/*
//runtime/slice.go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
*/

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println(s)
}

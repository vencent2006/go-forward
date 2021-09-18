/**
 * @Author: vincent
 * @Description:
 * @File:  slice_test
 * @Version: 1.0.0
 * @Date: 2021/8/18 09:30
 */

package slice_append

import (
	"fmt"
	"testing"
)

// case 1: modify by value
//-- outerSlice 1-- 0xc00000c0c0 [a a]   0xc00000c0e0 len:2 cap:2
//-- innerSlice 2-- 0xc00000c120 [a a]   0xc00000c0e0 len:2 cap:2
//-- innerSlice 3-- 0xc00000c120 [b b a]   0xc0000106c0 len:3 cap:4
//-- outerSlice 4-- 0xc00000c0c0 [a a]   0xc00000c0e0 len:2 cap:2
func modifySlice(innerSlice []string) {
	fmt.Printf("-- innerSlice 2-- %p %v   %p len:%d cap:%d\n", &innerSlice, innerSlice, &innerSlice[0], len(innerSlice), cap(innerSlice))
	innerSlice = append(innerSlice, "a")
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Printf("-- innerSlice 3-- %p %v   %p len:%d cap:%d\n", &innerSlice, innerSlice, &innerSlice[0], len(innerSlice), cap(innerSlice))
}

func TestSlice(t *testing.T) {
	outerSlice := []string{"a", "a"}
	fmt.Printf("-- outerSlice 1-- %p %v   %p len:%d cap:%d\n", &outerSlice, outerSlice, &outerSlice[0], len(outerSlice), cap(outerSlice))
	modifySlice(outerSlice)
	fmt.Printf("-- outerSlice 4-- %p %v   %p len:%d cap:%d\n", &outerSlice, outerSlice, &outerSlice[0], len(outerSlice), cap(outerSlice))
}

// case 2: modify by pointer
//[b b a]
//[b b a]
func modifySlice2(innerSlice *[]string) {
	*innerSlice = append(*innerSlice, "a")
	(*innerSlice)[0] = "b"
	(*innerSlice)[1] = "b"
	fmt.Println(*innerSlice)
}

func TestSlice2(t *testing.T) {
	outerSlice := []string{"a", "a"}
	modifySlice2(&outerSlice)
	fmt.Println(outerSlice)
}

// case 3
//[b b a]
//[b a]
func modifySlice3(innerSlice []string) {
	innerSlice[0] = "b"
	fmt.Printf("innerSlice[0]:%p\n", &innerSlice[0])
	innerSlice = append(innerSlice, "a")
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
	fmt.Printf("-- innerSlice 2-- %p %v   %p len:%d cap:%d\n", &innerSlice, innerSlice, &innerSlice[0], len(innerSlice), cap(innerSlice))
	fmt.Printf("innerSlice[1]:%p\n", &innerSlice[1])
}

func TestSlice3(t *testing.T) {
	outerSlice := []string{"a", "a"}
	modifySlice3(outerSlice)
	fmt.Println(outerSlice)
	fmt.Printf("-- outerSlice 4-- %p %v   %p len:%d cap:%d\n", &outerSlice, outerSlice, &outerSlice[0], len(outerSlice), cap(outerSlice))
	fmt.Printf("outerSlice[1]:%p\n", &outerSlice[1])
}

// case 4
//[b b a]
//[b b]
func modifySlice4(innerSlice []string) {
	innerSlice = append(innerSlice, "a")
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
	fmt.Printf("-- innerSlice 2-- %p %v   %p len:%d cap:%d\n", &innerSlice, innerSlice, &innerSlice[0], len(innerSlice), cap(innerSlice))
	fmt.Printf("innerSlice[1]:%p\n", &innerSlice[1])
}

func TestSlice4(t *testing.T) {
	outerSlice := make([]string, 0, 3)
	outerSlice = append(outerSlice, "a", "a")
	modifySlice4(outerSlice)
	fmt.Println(outerSlice)
	fmt.Printf("-- outerSlice 4-- %p %v   %p len:%d cap:%d\n", &outerSlice, outerSlice, &outerSlice[0], len(outerSlice), cap(outerSlice))
	fmt.Printf("outerSlice[1]:%p\n", &outerSlice[1])
}

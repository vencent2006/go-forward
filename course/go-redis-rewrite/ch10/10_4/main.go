package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "moody"
	st := reflect.TypeOf(s)
	fmt.Println(st)
	sv := reflect.ValueOf(s)
	fmt.Println(sv)
	s2 := sv.Interface().(string)
	fmt.Println(s2)
}

package main

import (
	"fmt"
	"unsafe"
)

type S1 struct {
	num1 int32
	num2 int32
}

type S2 struct {
	num1 int16
	num2 int32
}

func main() {
	fmt.Println(unsafe.Sizeof(S1{})) // 8
	fmt.Println(unsafe.Sizeof(S2{})) // 8 ，因为对齐
	fmt.Println("************")
	/*
	 Go的对齐系数：unsafe.Alignof
	*/
	fmt.Printf("bool size:%d align:%d\n", unsafe.Sizeof(bool(true)), unsafe.Alignof(bool(true)))
	fmt.Printf("byte size:%d align:%d\n", unsafe.Sizeof(byte(0)), unsafe.Alignof(byte(0)))
	fmt.Printf("int8 size:%d align:%d\n", unsafe.Sizeof(int8(0)), unsafe.Alignof(int8(0)))
	fmt.Printf("int16 size:%d align:%d\n", unsafe.Sizeof(int16(0)), unsafe.Alignof(int16(0)))
	fmt.Printf("int32 size:%d align:%d\n", unsafe.Sizeof(int32(0)), unsafe.Alignof(int32(0)))
	fmt.Printf("int64 size:%d align:%d\n", unsafe.Sizeof(int64(0)), unsafe.Alignof(int64(0)))
	fmt.Printf("string size:%d align:%d\n", unsafe.Sizeof(string("")), unsafe.Alignof(string("")))
}

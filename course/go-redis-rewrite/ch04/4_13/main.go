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

type Demo struct {
	a bool   // bool: 大小1，对齐系数1
	b string // string: 大小16，对齐系数8
	c int16  // int16：大小2，对齐系数2
}

type Demo1 struct {
	a bool   // bool: 大小1，对齐系数1
	c int16  // int16：大小2，对齐系数2
	b string // string: 大小16，对齐系数8
}

type Demo2 struct {
	a bool     // bool: 大小1，对齐系数1
	z struct{} // 空结构体单独出现时，地址为zerobase;空结构体出现在结构体中时，地址跟随前一个变量
	c int16    // int16：大小2，对齐系数2
	b string   // string: 大小16，对齐系数8
}

type Demo3 struct {
	a bool     // bool: 大小1，对齐系数1
	c int16    // int16：大小2，对齐系数2
	b string   // string: 大小16，对齐系数8
	z struct{} // 空结构体单独出现时，地址为zerobase;空结构体出现在结构体末尾时，需要补齐字长
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
	fmt.Printf("Demo size:%d align:%d\n", unsafe.Sizeof(Demo{}), unsafe.Alignof(Demo{}))
	fmt.Printf("Demo1 size:%d align:%d\n", unsafe.Sizeof(Demo1{}), unsafe.Alignof(Demo1{}))
	fmt.Printf("Demo2 size:%d align:%d\n", unsafe.Sizeof(Demo2{}), unsafe.Alignof(Demo2{}))
	fmt.Printf("Demo3 size:%d align:%d\n", unsafe.Sizeof(Demo3{}), unsafe.Alignof(Demo3{}))
}

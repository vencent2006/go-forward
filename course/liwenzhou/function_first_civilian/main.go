package main

import "fmt"

// Invoker 调用器接口
type Invoker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

func main() {
	var invoker Invoker
	s := new(Struct)
	invoker = s
	invoker.Call("hello")

	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	invoker.Call("hello")
}

package oop_inherit

import (
	"fmt"
	"testing"
)

type Animal struct {
}

func (a *Animal) SayHello() {
	fmt.Println("Hello!")
}

type Cat struct {
	*Animal
}

func TestPointer(t *testing.T) {
	cat := &Cat{&Animal{}}
	cat.SayHello()
}

func TestStruct(t *testing.T) {
	var cat Cat
	cat.SayHello()
}

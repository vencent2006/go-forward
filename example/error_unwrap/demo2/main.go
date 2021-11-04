package main

import (
	"errors"
	"fmt"
)

// Err1
type Err1 struct {
	Err error
}

func (e *Err1) Error() string { return "err1" }
func (e *Err1) Unwrap() error { return e.Err }

// Err2
type Err2 struct {
	Err error
}

func (e *Err2) Error() string { return "err2" }
func (e *Err2) Unwrap() error { return e.Err }

// Err3
type Err3 struct {
	Err error
}

func (e *Err3) Error() string { return "err3" }
func (e *Err3) Unwrap() error { return e.Err }

// 产生 error
func genErr() error {
	return &Err1{
		Err: &Err2{
			Err: &Err3{
				Err: nil,
			},
		},
	}
}

func main() {
	err := genErr()
	var err3 *Err3
	fmt.Println(errors.As(err, &err3)) // 第二个参数要求是 指针的地址
}

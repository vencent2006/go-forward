package main

import (
	"errors"
	"fmt"
	"reflect"
)

// Err1
type Err1 struct {
	Err error
}

func (e *Err1) Error() string { return "err1" }
func (e *Err1) Unwrap() error { return e.Err }
func (e *Err1) Is(other error) bool {
	v1 := reflect.ValueOf(e)
	v2 := reflect.ValueOf(other)
	// 如果是空指针直接返回
	if v1.IsNil() || v2.IsNil() {
		return false
	}
	// 取出指针指向的变量
	v1 = v1.Elem()
	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}
	return v1.IsValid() && v2.IsValid() && v1.Interface() == v2.Interface()
}

// Err2
type Err2 struct {
	Err error
}

func (e *Err2) Error() string { return "err2" }
func (e *Err2) Unwrap() error { return e.Err }
func (e *Err2) Is(other error) bool {
	v1 := reflect.ValueOf(e)
	v2 := reflect.ValueOf(other)
	// 如果是空指针直接返回
	if v1.IsNil() || v2.IsNil() {
		return false
	}
	// 取出指针指向的变量
	v1 = v1.Elem()
	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}
	return v1.IsValid() && v2.IsValid() && v1.Interface() == v2.Interface()
}

// Err3
type Err3 struct {
	Err error
}

func (e *Err3) Error() string { return "err3" }
func (e *Err3) Unwrap() error { return e.Err }
func (e *Err3) Is(other error) bool {
	v1 := reflect.ValueOf(e)
	v2 := reflect.ValueOf(other)
	// 如果是空指针直接返回
	if v1.IsNil() || v2.IsNil() {
		return false
	}
	// 取出指针指向的变量
	v1 = v1.Elem()
	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}
	return v1.IsValid() && v2.IsValid() && v1.Interface() == v2.Interface()
}

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
	err3 := &Err3{Err: nil}
	fmt.Println(errors.Is(err, err3))
}

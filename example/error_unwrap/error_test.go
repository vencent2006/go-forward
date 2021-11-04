/**
 * @Author: vincent
 * @Description:
 * @File:  error_test
 * @Version: 1.0.0
 * @Date: 2021/8/9 15:11
 */

package error_unwrap

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

// https://www.cnblogs.com/twoheads/p/12845404.html

// w
func TestW(t *testing.T) {
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err2: [%w]", err2)

	fmt.Println(err3)
}

// Unwrap
func TestUnwrap(t *testing.T) {
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err2: [%w]", err2)

	fmt.Println(errors.Unwrap(err3))
	fmt.Println(errors.Unwrap(err2))
}

// Is
func TestIs(t *testing.T) {
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err2: [%w]", err2)

	fmt.Println(errors.Is(err3, err2))
	fmt.Println(errors.Is(err3, err1))

}

// total
func TestAll(t *testing.T) {

	// New
	fmt.Println("-------New--------")
	err1 := errors.New("error1")
	err2 := errors.New("error2")

	// Is
	fmt.Println("-------Is--------")
	fmt.Println(errors.Is(err1, err2))
	fmt.Println(errors.Is(err1, errors.New("error1")))

	// As
	fmt.Println("-------As--------")
	fmt.Println(errors.As(err1, &err2)) // As的话，值相同就可以了；Is必须类型也得相同
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

	// Unwrap
	fmt.Println("-------Unwrap--------")
	e := errors.New("e")
	e1 := fmt.Errorf("e1: %w", e)
	e2 := fmt.Errorf("e2: %w", e1)
	fmt.Println(e2)
	fmt.Println(errors.Unwrap(e2))
	fmt.Println(e1)
	fmt.Println(errors.Unwrap(e1))

}

func TestUnwrap2(t *testing.T) {

	var e, first error

	first = errors.New("head")

	e = fmt.Errorf("第一层错误:%w", first)
	e = fmt.Errorf("第二层错误:%w", e)
	e = fmt.Errorf("第三层错误:%w", e)
	e = fmt.Errorf("第四层错误:%w", e)

	//解包错误
	e = errors.Unwrap(e)
	fmt.Printf("%s\n", e)

	//解包错误
	e = errors.Unwrap(e)
	fmt.Printf("%s\n", e)

	//解包错误
	e = errors.Unwrap(e)
	fmt.Printf("%s\n", e)

	//解包错误
	e = errors.Unwrap(e)
	fmt.Printf("%s\n", e)
}

func TestIsAsIs(t *testing.T) {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	fmt.Printf("err1 addr = %p, %+v\n", err1, err1)
	fmt.Printf("err2 addr = %p, %+v\n", err2, err2)
	fmt.Println()

	fmt.Println("--- TypeOf(err1) = ", reflect.TypeOf(err1))
	fmt.Println("--- TypeOf(err2) = ", reflect.TypeOf(err2))
	fmt.Println()

	fmt.Printf("---* err1 addr = %p, %+v\n", err1, err1)
	fmt.Printf("---* err2 addr = %p, %+v\n", err2, err2)
	fmt.Println()

	// todo: 这里Is返回的false，确实err1和err2的地址不一样
	fmt.Println("errors.Is(err1, err2) = ", errors.Is(err1, err2))
	fmt.Printf("---*- err1 addr = %p, %+v\n", err1, err1)
	fmt.Printf("---*- err2 addr = %p, %+v\n", err2, err2)
	fmt.Println()
	// todo: 这里As返回的是true，
	fmt.Println("errors.As(err1, err2) = ", errors.As(err1, &err2))
	fmt.Printf("---*-1 err1 addr = %p, %+v\n", err1, err1)
	fmt.Printf("---*-1 err2 addr = %p, %+v\n", err2, err2)
	fmt.Println()
	// todo: 这里Is返回的是true，因为err1的地址编程err2的地址，具体看errors.As的源码
	fmt.Println("errors.Is(err1, err2) = ", errors.Is(err1, err2))
	fmt.Printf("---*-1-2 err1 addr = %p, %+v\n", err1, err1)
	fmt.Printf("---*-1-2 err2 addr = %p, %+v\n", err2, err2)
	fmt.Println()

	fmt.Println("err1 == err2, is ", err1 == err2)

}

// As
type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}
func TestAs(t *testing.T) {

	var targetErr *ErrorString

	err := fmt.Errorf("new error:[%w]", &ErrorString{s: "target err"})
	err2 := fmt.Errorf("new error:[%w]", &ErrorString{s: "target err2222222222"})
	fmt.Println("errors.As is ", errors.As(err, &targetErr))
	fmt.Println("errors.Is is ", errors.Is(err, targetErr))
	// As 值相同就可以了， Is要求类型也要相同
	fmt.Println("errors.As(err, err2) is ", errors.As(err, &err2))
	fmt.Println("errors.Is(err, err2) is ", errors.Is(err, err2))

}

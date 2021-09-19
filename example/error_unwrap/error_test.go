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
	fmt.Println("errors.As is ", errors.As(err, &targetErr))
	fmt.Println("errors.Is is ", errors.Is(err, targetErr))

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
	fmt.Println(errors.As(err1, &err2)) // As的话，值相同就可以了；Is必须类型也得相同
	fmt.Println("-------As--------")
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

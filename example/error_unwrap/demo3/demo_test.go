package demo3

import (
	"errors"
	"fmt"
	"testing"
)

type MyError struct {
}

func (t *MyError) Error() string {
	return "Hello, it's my error"
}

func Test_errors(t *testing.T) {
	err := &MyError{}
	wrappedErr := fmt.Errorf("this is an wrapped error %w", err)

	if err == errors.Unwrap(wrappedErr) {
		fmt.Println("unwrapped")
	}

	if errors.Is(wrappedErr, err) {
		fmt.Println("wrapped is err")
	}

	copyErr := &MyError{}
	if errors.As(wrappedErr, &copyErr) {
		fmt.Println("convert error")
	}
}

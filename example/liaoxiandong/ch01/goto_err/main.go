package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	str := "aaa"
	err := genSomeErr(str)
	if err != nil {
		goto doExit
	}
	fmt.Println(str)
	str = "bbb"
	err = genSomeErr(str)
	if err != nil {
		goto doExit
	}
	fmt.Println(str)
	str = "ccc"
	err = genSomeErr(str)
	if err != nil {
		goto doExit
	}
	fmt.Println(str)

doExit:
	if err != nil {
		fmt.Println("err is", err)
	} else {
		fmt.Println("no err")
	}
}

func genSomeErr(errMsg string) error {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	log.Println("n = ", n)
	if n%2 == 0 {
		return errors.New(errMsg)
	} else {
		return nil
	}
}

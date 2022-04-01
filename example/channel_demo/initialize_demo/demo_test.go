/**
 * @Author: vincent
 * @Description:
 * @File:  demo_test
 * @Version: 1.0.0
 * @Date: 2022/4/1 09:22
 */

package initialize_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestNotInitialize(t *testing.T) {
	var c chan int
	var a int
	fmt.Println(c)
	go func() {
		fmt.Println("enter 11111")
		c <- 1
		fmt.Println("enter 333333")
	}()
	go func() {
		fmt.Println("enter 2222222")
		a = <-c
		fmt.Println(a)
		fmt.Println("enter 4444444")
	}()
	time.Sleep(1 * time.Second)
}

/**
 * @Author: vincent
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/4/8 20:13
 */

package main

import (
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	// trace prepare
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// my code here
	var total int
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				total += readNumber()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

//go:noinline
func readNumber() int {
	return rand.Intn(10)
}

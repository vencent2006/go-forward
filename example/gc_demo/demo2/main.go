package main

import (
	"os"
	"runtime/trace"
)

var cache = map[interface{}]interface{}{}
var ch = make(chan struct{})

func keepalloc() {
	for i := 0; i < 100; i++ {
		m := make([]byte, 1<<10)
		cache[i] = m
	}
}

func keepalloc2() {
	for i := 0; i < 100000; i++ {
		go func() {
			select {}
		}()
	}
}

func keepalloc3() {
	for i := 0; i < 100000; i++ {
		go func() {
			ch <- struct{}{}
		}()
	}
}

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	// test functions
	keepalloc()
	keepalloc2()
	keepalloc3()
}

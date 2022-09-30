package main

import "sync"

func a() {
	mu := sync.Mutex{}
	defer mu.Unlock()
	mu.Lock()
	// *****
}
func main() {
	a()
}

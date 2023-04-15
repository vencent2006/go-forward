package tests

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func Test_limit(t *testing.T) {
	var eg errgroup.Group
	eg.SetLimit(0)
	ch := make(chan int, 1)
	select {
	case ch <- func() int {
		eg.Go(func() error {
			return nil
		})
		return 1
	}():
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
	//eg.Go(func() error {
	//	fmt.Println("finish job")
	//	return nil
	//})
	err := eg.Wait()
	t.Log(err)
}

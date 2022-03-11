package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1000 * time.Millisecond

func main() {
	// 在指定deadline
	d := time.Now().Add(shortDuration)
	// 创建有指定deadline的context
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	// select handle event
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Second

func main() {
	// 创建截止时间
	d := time.Now().Add(shortDuration)
	// 创建有截止时间的context
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	// 使用select监听1s和有截止时间的context，谁先到来
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

package main

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/ratelimit"
)

func TestRateLimit(t *testing.T) {
	rl := ratelimit.New(100) // per second

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()              // 等待1/100秒，即10ms
		fmt.Println(i, now.Sub(prev)) // 打印距离上次操作(即Take)的时间间隔
		prev = now
	}
}

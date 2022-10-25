package main

import (
	"net/http"
	"time"

	tokenBucket "github.com/juju/ratelimit"

	"github.com/gin-gonic/gin"
	leakyBucket "go.uber.org/ratelimit" // 漏桶
)

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func heiHandler(c *gin.Context) {
	c.String(http.StatusOK, "ha")
}

// 基于漏桶的限流中间件1
func rateLimitLeakyBucket() func(c *gin.Context) {
	// 单位是1ns，1s=10的9次方纳秒
	rl := leakyBucket.New(10)
	return func(c *gin.Context) {
		// 取水滴
		prev := time.Now()
		now := rl.Take()
		delta := now.Sub(prev)
		if delta > time.Millisecond {

			//time.Sleep(delta) // 需要等这么长时间下一滴水才会到来
			c.String(http.StatusOK, "[leaky bucket] rate limit:"+delta.String())
			c.Abort()
			return
		}
		c.Next()
	}
}

// 基于令牌桶的限流中间件2
func rateLimitTokenBucket(fillInterval time.Duration, capacity int64) func(c *gin.Context) {
	rl := tokenBucket.NewBucket(fillInterval, capacity)
	return func(c *gin.Context) {
		//rateLimit
		//rl.Take() // 可以欠账
		//rl.TakeAvailable() // 不可以欠账
		var requiredCount int64 = 1
		if rl.TakeAvailable(requiredCount) == requiredCount {
			//	取到目标数量令牌
			c.Next()
			return
		}
		c.String(http.StatusOK, "[token bucket] rate limit... ")
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", rateLimitLeakyBucket(), pingHandler)
	r.GET("/hei", rateLimitTokenBucket(time.Second*2, 1), heiHandler)
	r.Run()
}

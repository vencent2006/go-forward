/**
 * @Author: vincent
 * @Description:
 * @File:  cost
 * @Version: 1.0.0
 * @Date: 2021/9/21 23:21
 */

package middleware

import (
	"go-examples/course/handwriting-web-inf/code_09/framework/gin"
	"log"
	"time"

	"go.uber.org/zap"
)

// cost机制，计算调用的时间成本
func Cost() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		// 记录开始时间
		start := time.Now()
		zap.L().Info("come here")

		// 使用next执行具体的业务逻辑
		c.Next()

		// 记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost: %v", c.Request.RequestURI, cost.Seconds())

	}
}

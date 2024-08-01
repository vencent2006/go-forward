package month

import (
	"fmt"
	"testing"
	"time"
)

func TestMonth(t *testing.T) {

	// 获取当前时间
	now := time.Now()

	// 获取当前月份
	month := now.Month()

	// 打印月份
	fmt.Printf("当前月份: %v\n", int(month))
}

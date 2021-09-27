/**
 * @Author: vincent
 * @Description:
 * @File:  max_test
 * @Version: 1.0.0
 * @Date: 2021/9/23 18:01
 */

package inline_func_demo

import "testing"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var Result int

func BenchmarkMax(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = max(-1, 1)
	}
	Result = r
}

/**
 * @Author: vincent
 * @Description:
 * @File:  slice_test
 * @Version: 1.0.0
 * @Date: 2021/11/13 17:29
 */

package slice_perf

import "testing"

//go:noinline
func callSlice(s []int) {

}

func BenchmarkCallSlice(b *testing.B) {
	s := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		callSlice(s)
	}
}

//go:noinline
func callArray(a [10000]int) {

}

func BenchmarkCallArray(b *testing.B) {
	var a [10000]int
	for i := 0; i < b.N; i++ {
		callArray(a)
	}
}

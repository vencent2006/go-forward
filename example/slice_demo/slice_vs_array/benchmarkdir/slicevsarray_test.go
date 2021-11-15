/**
 * @Author: vincent
 * @Description:
 * @File:  slicevsarray_benchmark
 * @Version: 1.0.0
 * @Date: 2021/11/11 20:33
 */

package benchmarkdir

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

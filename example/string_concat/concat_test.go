/**
 * @Author: vincent
 * @Description:
 * @File:  concat_test
 * @Version: 1.0.0
 * @Date: 2021/7/19 15:34
 */

package main

import (
	"fmt"
	"testing"
)

func BenchmarkPlusConcat(b *testing.B) {
	var main string
	for i := 0; i < b.N; i++ {
		main = ""
		main += "userid : " + "1"
		main += "location : " + "ab"
	}
}

func BenchmarkSprintf(b *testing.B) {
	var main string
	for i := 0; i < b.N; i++ {
		main = ""
		main += fmt.Sprintf("userid : %v", "1")
		main += fmt.Sprintf("location: %v", "ab")
	}
}

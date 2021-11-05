/**
 * @Author: vincent
 * @Description:
 * @File:  solution_test
 * @Version: 1.0.0
 * @Date: 2021/11/5 00:50
 */

package _354_russian_doll_envelopes

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func TestMaxEnvelopes(t *testing.T) {
	qs := []question{
		{
			para{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}},
			ans{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem ------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxEnvelopes(p.one))
	}
	fmt.Printf("\n\n\n")
}

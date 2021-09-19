/**
 * @Author: vincent
 * @Description:
 * @File:  solution_test
 * @Version: 1.0.0
 * @Date: 2021/9/19 22:09
 */

package _137_n_th_tribonacci_number

import "testing"

type para struct {
	n int
}

type ans struct {
	num int
}

type question struct {
	para
	ans
}

func TestTribonacci(t *testing.T) {
	qs := []question{
		{
			para{n: 3},
			ans{num: 2},
		},
		{
			para{n: 25},
			ans{num: 1389537},
		},
		{
			para{n: 4},
			ans{num: 4},
		},
	}

	for _, q := range qs {
		output := tribonacci(q.para.n)
		if output != q.ans.num {
			t.Logf("para:%v, ans:%v, output:%v", q.para, q.ans, output)
		}
	}
}

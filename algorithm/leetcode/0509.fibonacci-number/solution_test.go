/**
 * @Author: vincent
 * @Description:
 * @File:  solution_test
 * @Version: 1.0.0
 * @Date: 2021/9/19 21:08
 */

package _509_fibonacci_number

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

func TestFib(t *testing.T) {
	qs := []question{
		{
			para{n: 2},
			ans{num: 1},
		},
		{
			para{n: 3},
			ans{num: 2},
		},
		{
			para{n: 4},
			ans{num: 3},
		},
	}

	for _, q := range qs {
		output := fib(q.para.n)
		if output != q.ans.num {
			t.Logf("para:%v, ans:%v, output:%v", q.para, q.ans, output)
		}
	}
}

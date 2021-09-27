/**
 * @Author: vincent
 * @Description:
 * @File:  solution_test
 * @Version: 1.0.0
 * @Date: 2021/9/27 09:46
 */

package _018_4sum

import (
	"reflect"
	"sort"
	"testing"
)

type para struct {
	nums   []int
	target int
}

type ans struct {
	res [][]int
}

type question struct {
	para
	ans
}

func TestFourSum(t *testing.T) {

	qs := []question{
		{
			para{nums: []int{1, 0, -1, 0, -2, 2}, target: 0},
			ans{res: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		},
		{
			para{nums: []int{2, 2, 2, 2, 2}, target: 8},
			ans{res: [][]int{{2, 2, 2, 2}}},
		},
	}

	for _, q := range qs {
		output := fourSum(q.para.nums, q.para.target)

		if !isSame(output, q.ans.res) {
			t.Errorf("para:%v, ans:%v, output:%v", q.para, q.ans, output)
		}
	}
}

func isSame(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !isSame2(a[i], b[i]) {
			return false
		}
	}
	return true
}

func isSame2(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)
	return reflect.DeepEqual(a, b)
}

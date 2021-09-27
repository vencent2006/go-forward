/**
 * @Author: vincent
 * @Description:
 * @File:  solution_test
 * @Version: 1.0.0
 * @Date: 2021/9/27 09:40
 */

package _015_3sum

import (
	"reflect"
	"sort"
	"testing"
)

func TestSliceRange(t *testing.T) {
	s := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 0}}
	var res [][]int
	for _, item := range s {
		item = append(item, 88)
		res = append(res, item)
	}
	t.Logf("res is %+v", res)
}

type para struct {
	nums []int
}

type ans struct {
	res [][]int
}

type question struct {
	para
	ans
}

func TestThreeSum(t *testing.T) {

	qs := []question{
		{
			para{nums: []int{-1, 0, 1, 2, -1, -4}},
			ans{res: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		},
		{
			para{nums: []int{}},
			ans{res: [][]int{}},
		},
		{
			para{nums: []int{0}},
			ans{res: [][]int{}},
		},
	}

	for _, q := range qs {
		output := threeSum(q.para.nums)

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

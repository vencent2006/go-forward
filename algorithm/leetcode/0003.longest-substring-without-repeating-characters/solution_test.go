/**
 * @Author: vincent
 * @Description:
 * @File:  solution_test
 * @Version: 1.0.0
 * @Date: 2021/9/29 09:23
 */

package _003_longest_substring_without_repeating_characters

import "testing"

type para struct {
	one string
}

type ans struct {
	one int
}

type question struct {
	para
	ans
}

func TestLengthOfLongestSubstring(t *testing.T) {
	qs := []question{
		{
			para{one: "abcabcbb"},
			ans{one: 3},
		},
		{
			para{one: "bbbbb"},
			ans{one: 1},
		},
		{
			para{one: "pwwkew"},
			ans{one: 3},
		},
		{
			para{one: ""},
			ans{one: 0},
		},
	}

	for _, q := range qs {
		output := lengthOfLongestSubstring(q.para.one)
		if output != q.ans.one {
			t.Errorf("para:%v, ans:%v, output:%v", q.para.one, q.ans.one, output)
		}
	}
}

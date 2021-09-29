/**
 * @Author: vincent
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2021/9/29 09:17
 */

package _003_longest_substring_without_repeating_characters

// 核心思想: 滑动窗口
func lengthOfLongestSubstring(s string) int {
	window := make(map[uint8]int)
	left, right := 0, 0
	res := 0
	length := len(s)
	for right < length {
		c := s[right]
		right++

		// 进行窗口内数据的一系列更新
		window[c]++
		// 判断左侧窗口是否需要收缩
		for window[c] > 1 {
			// 有重复
			d := s[left]
			left++

			// 进行窗口内数据的一系列更新
			window[d]--
		}
		// 在这里更新答案
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

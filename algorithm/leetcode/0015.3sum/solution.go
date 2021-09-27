/**
 * @Author: vincent
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2021/9/27 09:12
 */

package _015_3sum

import (
	"sort"
)

func twoSumTarget(nums []int, start int, target int) [][]int {
	l, h := start, len(nums)-1
	var res [][]int
	var sum, left, right int
	for l < h {
		sum = nums[l] + nums[h]
		left, right = nums[l], nums[h]
		if sum < target {
			// l向后找
			for l < h && nums[l] == left {
				l++
			}
		} else if sum > target {
			// h向前找
			for l < h && nums[h] == right {
				h--
			}
		} else {
			// sum == target
			res = append(res, []int{left, right})
			for l < h && nums[l] == left {
				// l向后找
				l++
			}
			for l < h && nums[h] == right {
				h--
			}
		}
	}

	return res
}

func threeSum(nums []int) [][]int {
	// sort first
	sort.Ints(nums)

	n := len(nums)
	var res [][]int
	target := 0
	for i := 0; i < n; i++ {
		// find two-sum tuple(元组)
		tuples := twoSumTarget(nums, i+1, target-nums[i])
		for _, tuple := range tuples {
			tuple = append(tuple, nums[i])
			res = append(res, tuple)
		}

		// 等值向后过滤
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

/**
 * @Author: vincent
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2021/9/28 09:04
 */

package _000_n_sum

func nSumTarget(nums []int, n, start, target int) [][]int {
	// 数组长度
	sz := len(nums)
	var res [][]int
	// n至少是2sum，且数组大小不应该小于n
	if n < 2 || sz < n {
		return res
	}

	// 2sum 是base case
	if n == 2 {
		lo, hi := start, sz-1
		for lo < hi {
			// 先求和
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				// lo右挪
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				// hi左挪
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else {
				res = append(res, []int{left, right})
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else {
		// n>2时，递归计算(n-1)Sum的结果
		for i := start; i < sz; i++ {
			subSums := nSumTarget(nums, n-1, i+1, target-nums[i])
			for _, subSum := range subSums {
				subSum = append(subSum, nums[i])
				res = append(res, subSum)
			}
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}

/**
 * @Author: vincent
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2021/11/3 00:42
 */

package _300_Longest_Increasing_Subsequence

func lengthOfLIS(nums []int) int {
	// dp[i] 表示以nums[i]结尾的最长递增子序列的长度
	dp, res := make([]int, len(nums)), 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLIS2(nums []int) int {
	// dp[i] 表示以nums[i]结尾的最长递增子序列的长度
	dp, res := make([]int, len(nums)), 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

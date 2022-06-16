/**
 * @Author: vincent
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2021/11/5 00:43
 */

package _354_russian_doll_envelopes

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	dp := make([]int, len(envelopes))
	var res int
	for i := range envelopes {
		dp[i] = 1
		for j := 0; j <= i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			res = max(res, dp[i])
		}
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

func maxEnvelopes2(envelops [][]int) int {
	sort.Slice(envelops, func(i, j int) bool {
		return envelops[i][0] < envelops[j][0] || (envelops[i][0] == envelops[j][0] && envelops[i][1] > envelops[j][1])
	})

	dp := make([]int, len(envelops))
	var res int
	for i := range envelops {
		dp[i] = 1
		for j := 0; j <= i; j++ {
			if envelops[j][1] < envelops[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			res = max(res, dp[i])
		}
	}
	return res
}

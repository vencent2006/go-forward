/**
 * @Author: vincent
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2021/9/19 00:03
 */

package _001_two_sum

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for index, num := range nums {
		if i, ok := hashMap[target-num]; ok {
			return []int{i, index}
		}
		// 存的是下标
		hashMap[num] = index
	}
	return []int{}
}

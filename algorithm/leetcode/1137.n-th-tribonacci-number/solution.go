/**
 * @Author: vincent
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2021/9/19 21:59
 */

package _137_n_th_tribonacci_number

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	pre1, pre2, cur, sum := 0, 1, 1, 0
	for i := 3; i <= n; i++ {
		sum = pre1 + pre2 + cur
		// next iter
		pre1 = pre2
		pre2 = cur
		cur = sum
	}

	return cur
}

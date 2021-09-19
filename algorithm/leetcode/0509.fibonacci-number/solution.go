/**
 * @Author: vincent
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2021/9/19 21:02
 */

package _509_fibonacci_number

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 2 || n == 1 {
		return 1
	}

	// pre 前一个数值，cur 当前数值
	pre, cur, sum := 1, 1, 0
	for i := 3; i <= n; i++ {
		// sum是为了记录而存在
		sum = pre + cur
		// 迭代下一次的值
		pre = cur
		cur = sum
	}

	// 返回当前值
	return cur
}

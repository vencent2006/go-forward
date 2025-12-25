#include <iostream>
#include <algorithm>

using namespace std;

int main() {
    int n, a, b;
    cin >> n >> a >> b;

    // 对角线编号
    int d = a + n - b;

    long long result;

    if (d <= n) {
        // 前d-1条对角线的元素总数
        long long sum_before = 1LL * d * (d - 1) / 2;

        if (d % 2 == 1) {
            // 奇数对角线：从右上向左下填充（在矩阵中是从上到下）
            result = sum_before + (n - b + 1);
        } else {
            // 偶数对角线：从左下向右上填充（在矩阵中是从下到上）
            result = sum_before + a;
        }
    } else {
        // d > n 的情况
        // 前n条对角线的元素总数
        long long sum_before = 1LL * n * (n + 1) / 2;

        // 第n+1到第d-1条对角线的元素总数
        int m = d - n - 1;
        // 这些对角线的长度：n-1, n-2, ..., n-m
        // 求和公式：m * (2n - m - 1) / 2
        sum_before += 1LL * m * (2 * n - m - 1) / 2;

        // 当前对角线上的位置
        if (d % 2 == 1) {
            // 奇数对角线：从右上向左下填充
            result = sum_before + (a - (d - n));
        } else {
            // 偶数对角线：从左下向右上填充
            result = sum_before + (n - b + 1);
        }
    }

    cout << result << endl;

    return 0;
}
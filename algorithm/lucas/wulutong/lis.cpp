#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

// 计算最大公约数
int gcd(int a, int b) {
    while (b) {
        int temp = b;
        b = a % b;
        a = temp;
    }
    return a;
}

int main() {
    int n;
    cin >> n;

    vector<int> nums(n);
    for (int i = 0; i < n; i++) {
        cin >> nums[i];
    }

    // dp[i]：以第i个元素结尾的最长合法we子序列长度
    vector<int> dp(n, 1);  // 每个元素自身至少长度为1

    // 最大值
    int maxLen = 1;

    // 动态规划填表
    for (int i = 0; i < n; i++) {
        dp[i] = 1;  // 至少包含自身

        // 检查所有前面的元素
        for (int j = 0; j < i; j++) {
            // 如果nums[j]和nums[i]不互质，且dp[j]+1能使得dp[i]更大
            if (gcd(nums[j], nums[i]) > 1) {
                dp[i] = max(dp[i], dp[j] + 1);
            }
        }

        // 更新全局最大值
        maxLen = max(maxLen, dp[i]);
    }

    cout << maxLen << endl;

    return 0;
}
// 1. 首先检查糖果总数能否被n整除（平均数必须是整数）
// 2. 计算平均数
// 3. 计算前缀和数组，表示前i个小朋友的总糖果数
// 4. 前i个小朋友的理想总糖果数应该是 （i*平均数）
// 5. 实际前缀和 减 理想前缀和
// 6. 最大差值就是所需的最少轮数
#include <iostream>
using namespace std;
int a[10005], sum[10005];

int main() {
    int n;
    cin >> n;
    long long cnt = 0;
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
        cnt += a[i];
    }
    if (cnt % n != 0) {
        cout << -1;
        return 0;
    }
    int avg = cnt / n;
    for (int i = 1; i <= n; i++) {
        sum[i] = sum[i - 1] + a[i];
    }

    int res = 0;
    for (int i = 1; i <= n; i++) {
        res = max(res, abs(sum[i] - avg * i));
    }
    cout << res;
    return 0;
}

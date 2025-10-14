#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

struct Item {
    int v;  // 价格
    int w;  // 价值
    int s;  // 最大数量
};

int main() {
    int n, m;
    cin >> n >> m;

    vector<Item> items(n);
    for (int i = 0; i < n; i++) {
        cin >> items[i].v >> items[i].w >> items[i].s;
    }

    // dp[j] 表示花费j元能获得的最大价值
    vector<int> dp(m + 1, 0);

    // 多重背包
    for (int i = 0; i < n; i++) {
        int v = items[i].v;
        int w = items[i].w;
        int s = items[i].s;

        // 对于每种物品，考虑购买0到s个的情况
        for (int j = m; j >= v; j--) {
            for (int k = 1; k <= s && k * v <= j; k++) {
                dp[j] = max(dp[j], dp[j - k * v] + k * w);
            }
        }
    }

    cout << dp[m] << endl;

    return 0;
}
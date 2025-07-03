#include <iostream>
using namespace std;
// https://htoj.com.cn/cpp/oj/problem/detail?pid=22386408740480&cid=22394268493184&gid=22185991356800
int main() {
    int nums[100001] = {};
    int n;
    cin >> n;
    int sum = 0;
    for (int i = 1; i <= n; i++) {
        cin >> nums[i];
        sum += nums[i];
    }
    int left = 0;
    int min = INT_MAX;
    for (int i = 1; i <= n; i++) {
        left += nums[i];
        int diff = (sum - left) - left;
        diff = diff > 0 ? diff : -diff;
        if (diff < min) {
            min = diff;
        }
    }
    cout << min;
    return 0;
}

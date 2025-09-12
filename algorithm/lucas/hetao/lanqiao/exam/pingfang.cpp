#include <iostream>
#include <vector>
#include <map>
#include <cmath>

using namespace std;

int getSquareFree(int x) {
    if (x == 0) return 0;
    int res = 1;
    int n = x;
    for (int i = 2; i * i <= n; i++) {
        int cnt = 0;
        while (x % i == 0) {
            cnt++;
            x /= i;
        }
        if (cnt % 2 == 1) {
            res *= i;
        }
    }
    if (x > 1) {
        res *= x;
    }
    return res;
}

int main() {
    int n;
    cin >> n;
    vector<int> A(n);
    for (int i = 0; i < n; i++) {
        cin >> A[i];
    }

    map<int, int> group;
    int zero_count = 0;

    for (int num: A) {
        int s = getSquareFree(num);
        if (s == 0) {
            zero_count++;
        } else {
            group[s]++;
        }
    }

    int max_group_size = 0;
    for (auto it: group) {
        if (it.second > max_group_size) {
            max_group_size = it.second;
        }
    }

    int ans = max_group_size;
    if (zero_count > 0) {
        ans = max(ans, zero_count);
        ans = max(ans, max_group_size + zero_count);
    }

    cout << ans << endl;

    return 0;
}
#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;
    if (n % 2 == 1) {
        cout << -1;
        return 0;
    }

    // 这个输出的次序是从小到大的，不满足题目要求
    for (int i = 1; i <= 31; i++) {
        int t = 1 << i;
        if (t > n) {
            break;
        }
        if (t & n) {
            cout << t << " ";
        }
    }
    return 0;
}

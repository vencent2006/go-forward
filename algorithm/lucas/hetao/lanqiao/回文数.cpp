#include <iostream>
using namespace std;

bool isPalindromeWith7(int n) {
    int m = 0;
    int t = n;
    bool isHas7 = false;
    while (t > 0) {
        if (t % 10 == 7) {
            isHas7 = true;
        }
        m = m * 10 + t % 10;
        t /= 10;
    }
    return m == n && isHas7;
}

int main() {
    int n, m;
    cin >> n >> m;
    for (int i = n; i <= m; i++) {
        if (isPalindromeWith7(i)) {
            cout << i << " ";
        }
    }
    return 0;
}

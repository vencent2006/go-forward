#include <iostream>
using namespace std;

int main() {
    int n, m, cnt = 0;
    cin >> n >> m;
    for (int i = n; i <= m; i++) {
        int j = i;
        while (j != 0) {
            int t = j % 10;
            j /= 10;
            if (t == 7) {
                cnt++;
                break;
            }
        }
    }
    cout << cnt;
    return 0;
}

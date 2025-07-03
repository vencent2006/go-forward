#include <iostream>
using namespace std;

int main() {
    int a[1001] = {};
    int n;
    cin >> n;
    for (int i = 0; i < n; i++) {
        cin >> a[i];
    }
    //up[i] 表示以第 i 个元素结尾，并且最后一个差值为正的最长摆动子序列长度。
    //down[i] 表示以第 i 个元素结尾，并且最后一个差值为负的最长摆动子序列长度。
    int up = 1, down = 1; // up 表示当前位置是上升的最长子序列长度，down 表示当前位置是下降的最长子序列长度
    for (int i = 1; i < n; i++) {
        if (a[i] > a[i - 1]) {
            up = down + 1;
        } else if (a[i] < a[i - 1]) {
            down = up + 1;
        }
    }
    cout << max(up, down) << endl;

    return 0;
}

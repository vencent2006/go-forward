#include <iostream>
#include <cstring>
using namespace std;
const int N = 2000;
// 高精度数加法 前导零
int main() {
    char a[N], b[N];
    int a1[N], b1[N], c1[N];
    cin >> a >> b;
    int la = strlen(a);
    int lb = strlen(b);
    int lc = la > lb ? la : lb;

    for (int i = 0; i < la; i++) {
        a1[la - i - 1] = a[i] - '0';
    }

    for (int i = 0; i < lb; i++) {
        b1[lb - i - 1] = b[i] - '0';
    }

    int carry = 0;
    for (int i = 0; i < lc; i++) {
        c1[i] = a1[i] + b1[i] + carry;
        carry = c1[i] / 10;
        c1[i] %= 10;
    }
    if (carry > 0) {
        c1[lc] = carry;
        lc++;
    }
    // if (carry) {
    //     c1[lc] = carry;
    // } else {
    //     lc--;
    // }
    int k; // k 记录第一个非0数字下标
    // 从后往前记录第一个非0数字
    for (k = lc; k > 0; k--) {
        if (c1[k] != 0) break;
    }
    for (int i = k; i >= 0; i--) {
        cout << c1[i];
    }
    return 0;
}

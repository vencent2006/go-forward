#include <iostream>
#include <cstring>
#include <algorithm>
using namespace std;
char a[2001], b[2001], c[2001]; // 输入的字符串
int a1[2001], b1[2001], c1[2001]; // 把输入的字符串，反转存入数组，a1[0]是个位
int d1[5001]; // 加和

// 作业2811 高精度数加法 不同位数有进位 三个整数相加
int main() {
    cin >> a >> b >> c;
    int la = strlen(a), lb = strlen(b), lc = strlen(c);
    for (int i = 0; i < la; i++) {
        a1[la - i - 1] = a[i] - '0';
    }

    for (int i = 0; i < lb; i++) {
        b1[lb - i - 1] = b[i] - '0';
    }

    for (int i = 0; i < lc; i++) {
        c1[lc - i - 1] = c[i] - '0';
    }

    int ld = max(max(la, lb), lc);
    int carry = 0;
    for (int i = 0; i < ld; i++) {
        d1[i] = a1[i] + b1[i] + c1[i] + carry;
        carry = d1[i] / 10;
        d1[i] %= 10;
    }
    if (carry > 0) {
        c1[lc] = carry;
        ld++;
    }
    for (int i = ld - 1; i >= 0; i--) {
        cout << d1[i];
    }
    return 0;
}

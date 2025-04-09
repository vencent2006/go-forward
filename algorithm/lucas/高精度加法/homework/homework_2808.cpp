#include <iostream>
#include <cstring>
using namespace std;
char a[2001], b[2001];
int a1[2001], b1[2001], c1[2001];

// 高精度数加法 不同位数有进位
int main() {
    cin >> a >> b;
    int la = strlen(a), lb = strlen(b);
    for (int i = 0; i < la; i++) {
        a1[la - i - 1] = a[i] - '0';
    }


    for (int i = 0; i < lb; i++) {
        b1[lb - i - 1] = b[i] - '0';
    }

    int lc = la > lb ? la : lb;
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
    for (int i = lc - 1; i >= 0; i--) {
        cout << c1[i];
    }
    return 0;
}

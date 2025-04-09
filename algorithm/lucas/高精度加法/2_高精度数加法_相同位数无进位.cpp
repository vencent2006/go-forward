#include <iostream>
#include <cstring>
using namespace std;
char a[2001], b[2001];
int a1[2001], b1[2001], c1[2001];

// 高精度数加法 相同位数无进位
int main() {
    cin >> a >> b;
    const int la = strlen(a);
    for (int i = 0; i < la; i++) {
        a1[i] = a[i] - '0';
        b1[i] = b[i] - '0';
        c1[i] = a1[i] + b1[i];
    }
    for (int i = 0; i < la; i++) {
        cout << c1[i];
    }
    return 0;
}

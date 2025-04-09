#include <iostream>
#include <cstring>
using namespace std;
char a[2001];
int a1[2001];
// 高精度数存储
int main() {
    cin >> a;
    const int len = strlen(a);
    for (int i = 0; i < len; i++) {
        a1[i] = a[i] - '0';
    }
    for (int i = 0; i < len; i++) {
        cout << a1[i];
    }
    return 0;
}

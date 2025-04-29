#include <iostream>
//#include <bits/stdc++.h>
using namespace std;

int main() {
    int a, b, ans;
    cin >> a >> b; //35 94
    for (int i = 1; i < a; i++) {
        // i 是鸡， j 是兔
        int j = a - i;
        if (2 * i + 4 * j == b) {
            ans = i - j;
            break;
        }
    }
    cout << ans << endl;
    return 0;
}

#include <iostream>
using namespace std;

int main() {
    int p, q;
    cin >> p >> q;

    int A = 0, B = 0; // A和B的按钮得分，1表示按正确了，0表示没按正确
    if (p == 0) {
        A = 1;
    } else {
        A = 0;
    }

    if (q == 0) {
        B = 1;
    } else {
        B = 0;
    }


    for (int i = 1; i <= 4; i++) {
        int a[5] = {};
        a[i] = 1; // 假设是正确的按钮
        if (a[1] + a[3] == A && a[3] + a[4] == B) {
            cout << i << endl;
            return 0;
        }
    }


    return 0;
}

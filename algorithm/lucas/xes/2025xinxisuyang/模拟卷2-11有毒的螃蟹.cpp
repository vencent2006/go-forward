#include <iostream>
using namespace std;

int main() {
    int a, b;
    cin >> a >> b;
    // 0代表食物中毒，1代表吃了没事
    if (a == 0 && b == 0) {
        cout << 1;
    } else if (a == 0 && b == 1) {
        cout << 2;
    } else if (a == 1 && b == 0) {
        cout << 3;
    } else {
        cout << 4;
    }
    return 0;
}

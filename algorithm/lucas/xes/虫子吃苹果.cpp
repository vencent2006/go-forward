#include <iostream>
using namespace std;

int main() {
    int n, x, y;
    cin >> n >> x >> y;
    // int eaten = y / x;
    // if (y > eaten * x) {
    //     // 如果y不能被x整除, 则需要多浪费一个苹果
    //     eaten++;
    // }
    int eaten = (y + x - 1) / x;
    if (n >= eaten) {
        cout << n - eaten << endl;
    } else {
        cout << 0 << endl;
    }
    return 0;
}

#include <iostream>
using namespace std;

int f(int x) {
    if (x % 2 == 1) {
        return x * 2 + 3;
    } else {
        return x / 2 + 7;
    }
}

int main() {
    int a, b, c;
    cin >> a >> b >> c;
    cout << f(a) + f(b) + f(c) << endl;
    return 0;
}

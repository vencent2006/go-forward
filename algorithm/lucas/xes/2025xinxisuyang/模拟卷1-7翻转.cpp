#include <iostream>
using namespace std;

int reverse(int num) {
    // 123 321
    // 700 7
    int sum = 0;
    while (num > 0) {
        sum = sum * 10 + num % 10;
        num = num / 10;
    }
    return sum;
}

int main() {
    int n;
    cin >> n;
    for (int i = 0; i < n; i++) {
        int num;
        cin >> num;
        cout << reverse(num) << " ";
    }
    return 0;
}

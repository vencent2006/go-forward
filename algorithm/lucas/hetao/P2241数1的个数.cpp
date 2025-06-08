#include <iostream>
using namespace std;

int countOne(int num) {
    int count = 0;
    while (num > 0) {
        if (num % 10 == 1) {
            count++;
        }
        num /= 10;
    }
    return count;
}

int main() {
    int n;
    cin >> n;
    int sum = 0;
    for (int i = 1; i <= n; i++) {
        sum += countOne(i);
    }
    cout << sum << endl;
    return 0;
}

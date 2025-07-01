#include <iostream>
using namespace std;

bool isLucky(int n) {
    if (n % 2 == 1) {
        return false;
    }
    int sum = 0;
    while (n != 0) {
        sum += n % 10;
        n = n / 10;
    }
    return sum % 2 == 0;
}

int main() {
    int n;
    cin >> n;
    int count = 0;
    for (int i = 1; i <= n; i++) {
        if (isLucky(i)) {
            count++;
        }
    }
    cout << count << endl;
    return 0;
}

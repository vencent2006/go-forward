#include <iostream>

using namespace std;

bool isReap(int year) {
    if (year % 4 == 0 && year % 100 != 0) {
        return true;
    }
    if (year % 400 == 0) {
        return true;
    }
    return false;
}

int main() {
    int a, b;
    cin >> a >> b;
    int sum = 0;
    for (int i = a + 1; i < b; i++) {
        if (isReap(i)) {
            sum += i;
        }
    }
    cout << sum << endl;
    return 0;
}

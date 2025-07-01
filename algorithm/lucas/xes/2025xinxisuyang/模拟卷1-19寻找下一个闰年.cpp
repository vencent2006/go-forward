#include <iostream>
using namespace std;

bool isLeapYear(int year) {
    if (year % 4 == 0 && year % 100 != 0 || year % 400 == 0) {
        return true;
    } else {
        return false;
    }
}

int main() {
    int n;
    cin >> n;
    n++;
    while (true) {
        if (isLeapYear(n)) {
            cout << n << endl;
            break;
        }
        n++;
    }
    return 0;
}

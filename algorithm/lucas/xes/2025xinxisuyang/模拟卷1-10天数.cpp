#include <iostream>
using namespace std;

bool isRunNian(int year) {
    if (year % 4 == 0 && year % 100 != 0 || year % 400 == 0) {
        return true;
    }
    return false;
}

int main() {
    int year, month;
    cin >> year >> month;
    int daysArr[12] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
    if (isRunNian(year)) {
        daysArr[1] = 29;
    }
    int days = daysArr[month - 1];
    if (days == 0) {
        // 偶数
        cout << days / 2 << " " << days / 2 << endl;
    } else {
        // 奇数
        cout << days / 2 + 1 << " " << days / 2 << endl;
    }
    return 0;
}

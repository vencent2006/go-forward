#include <iostream>

using namespace std;

bool isLeapYear(int y) {
    if (y % 4 == 0 && y % 100 != 0 || y % 400 == 0) {
        return true;
    }
    return false;
}

int getDaysInMonth(int y, int m) {
    if (m == 2) {
        if (y % 4 == 0 && y % 100 != 0 || y % 400 == 0) {
            return 29;
        }
        return 28;
    }
    if (m == 4 || m == 6 || m == 9 || m == 11) {
        return 30;
    }
    return 31;
}

int main() {
    int y, m, d, h, k;
    cin >> y >> m >> d >> h >> k;
    int totalHours = h + k;
    int extraDays = totalHours / 24;
    h = totalHours % 24;
    d += extraDays;

    int daysInMonth = getDaysInMonth(y, m);
    if (d > daysInMonth) {
        d -= daysInMonth;
        m++;
        if (m > 12) {
            m = 1;
            y++;
        }
    }

    cout << y << " " << m << " " << d << " " << h << endl;


    return 0;
}

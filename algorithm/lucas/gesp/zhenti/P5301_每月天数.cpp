#include <iostream>

using namespace std;

int days(int y, int m) {
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
    int y, m;
    cin >> y >> m;
    cout << days(y, m) << endl;
    return 0;
}

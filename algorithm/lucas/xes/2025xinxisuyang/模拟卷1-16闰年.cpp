#include <iostream>
using namespace std;

bool isRunNian(int year) {
    if (year % 4 == 0 && year % 100 != 0 || year % 400 == 0) {
        return true;
    }
    return false;
}

int main() {
    int start, end;
    cin >> start >> end;
    for (int i = start; i <= end; i++) {
        if (isRunNian(i)) {
            cout << "YES" << endl;
            return 0;
        }
    }
    cout << "NO" << endl;

    return 0;
}

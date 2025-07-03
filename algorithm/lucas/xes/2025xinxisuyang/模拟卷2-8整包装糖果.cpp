#include <iostream>
using namespace std;

int main() {
    int x, y;
    cin >> x >> y;
    if (y % x == 0) {
        cout << "OK 0" << endl;
    } else {
        cout << "ADD " << x - (y % x) << endl;
    }
    return 0;
}

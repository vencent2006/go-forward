#include <iostream>

using namespace std;

int main() {
    int a, b, c;
    cin >> a >> b >> c;
    if (a + b > c || a + c > b || b + c > a) {
        if (a * a + b * b == c * c || a * a + c * c == b * b || b * b + c * c == a * a) {
            cout << "yes" << endl;
            return 0;
        }
        cout << "no" << endl;
    } else {
        cout << "not a triangle" << endl;
    }
    return 0;
}

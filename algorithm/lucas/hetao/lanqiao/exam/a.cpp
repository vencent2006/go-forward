#include <iostream>

using namespace std;

int func(int y) {
    y -= 5;
    cout << "X";
    return 0;
}

int main() {
    int x = 10, y = 5;
    if (x > y || func(y))
        cout << y;

    return 0;
}
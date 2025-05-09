#include <iostream>
using namespace std;

int main() {
    int n, part1 = 0, part2 = 0, part3 = 0;
    string isJiaji;
    cin >> n >> isJiaji;
    if (n > 0) {
        part1 = 8;
    }
    if (n > 1000) {
        part2 = (n - 1000 + (500 - 1)) / 500 * 4;
    }
    if (isJiaji == "y") {
        part3 = 5;
    }

    cout << part1 + part2 + part3;

    return 0;
}

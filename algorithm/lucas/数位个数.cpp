#include <iostream>
using namespace std;

int main() {
    long long n;
    cin >> n;
    int sum = 0;
    if (n == 0) { sum++; }
    while (n > 0) {
        n /= 10;
        sum++;
    }
    cout << sum << endl;

    return 0;
}

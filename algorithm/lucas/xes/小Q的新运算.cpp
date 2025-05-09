#include <iostream>
using namespace std;

int main() {
    long long a, b;
    cin >> a >> b;
    long long ans = a * b * (b + 1) / 2;
    cout << ans;
    return 0;
}

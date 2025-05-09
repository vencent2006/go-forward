#include <iostream>
using namespace std;

int main() {
    int n, from, to;
    cin >> n >> from >> to;
    int sum = 0;
    for (int i = from; i <= to; i++) {
        if (i % 10 == n || i % n == 0) {
            sum += i;
        }
    }
    cout << sum << endl;
    return 0;
}

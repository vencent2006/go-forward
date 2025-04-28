#include <iostream>
using namespace std;

int main() {
    int n, num;
    cin >> n;
    int sum = 0;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            cin >> num;
            if (i == 0 || i == n - 1 || j == 0 || j == n - 1) {
                sum += num;
            }
        }
    }

    cout << sum << endl;
    return 0;
}

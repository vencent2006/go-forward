#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;
    for (int i = 1; i <= n; i++) {
        int num;
        cin >> num;
        if (i == n) {
            cout << num << endl;
        }
    }
    return 0;
}

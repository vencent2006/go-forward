#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;
    int jishu = 0, oushu = 0;
    for (int i = 1; i <= n; i++) {
        int number;
        cin >> number;
        if (number % 2 > 0) {
            jishu++;
        } else {
            oushu++;
        }
    }
    cout << jishu << " " << oushu << endl;
    return 0;
}

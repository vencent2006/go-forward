#include <iostream>
using namespace std;

int main() {
    int a, b;
    cin >> a >> b;
    int cnt = 0;
    for (int i = a; i <= b; i++) {
        int j = i;
        while (j > 0) {
            if (j % 10 == 7) {
                cnt++;
                break;
            }
            j /= 10;
        }
    }
    cout << cnt << endl;
    return 0;
}

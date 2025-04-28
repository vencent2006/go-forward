#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;
    int cntJishu = 0;
    for (int i = 0; i < n; i++) {
        int num;
        cin >> num;
        if (num % 2 == 1) {
            cntJishu++;
        }
    }
    cout << cntJishu << " " << n - cntJishu << endl;
    return 0;
}

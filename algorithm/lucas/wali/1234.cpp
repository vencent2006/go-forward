#include <iostream>
using namespace std;
int a[1501];

int main() {
    int n;
    cin >> n;
    int m, sum = 0;
    for (int i = 1; i <= n; i++) {
        cin >> m;
        a[m]++;
        sum += m;
    }
    for (int i = 1500; i > 0; i--) {
        for (int j = 0; j < a[i]; j++) {
            cout << i << " ";
        }
    }
    cout << endl << sum << endl;
    return 0;
}

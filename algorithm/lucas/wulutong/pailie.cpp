#include <iostream>

using namespace std;
int a[11][10001];

int main() {
    int n, k;
    cin >> n >> k;

    for (int i = 1; i <= n; i++) {
        cin >> a[0][i];
    }
    for (int j = 1; j <= k; j++) {
        a[j][1] = (a[j - 1][1] + a[j - 1][2]) / 2;
        a[j][n] = (a[j - 1][n - 1] + a[j - 1][n]) / 2;

        for (int i = 2; i < n; i++) {
            a[j][i] = (a[j - 1][i - 1] + a[j - 1][i] + a[j - 1][i + 1]) / 3;
        }


    }
    for (int i = 1; i <= n; i++) {
        cout << a[k][i] << " ";
    }
    cout << endl;

    return 0;
}
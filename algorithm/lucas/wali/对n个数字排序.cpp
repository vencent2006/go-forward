#include <iostream>
using namespace std;
int a[1001];

int main() {
    int n;
    cin >> n;
    int m;
    for (int i = 1; i <= n; i++) {
        cin >> m;
        a[m]++;
    }
    for (int i = 1; i <= 1000; i++) {
        for (int j = 0; j < a[i]; j++) {
            cout << i << " ";
        }
    }
    return 0;
}

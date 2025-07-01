#include <iostream>
using namespace std;
int a[10];

int main() {
    int n;
    cin >> n;
    string m;
    for (int i = 0; i < n; i++) {
        cin >> m;
        for (int j = 0; j < m.size(); j++) {
            a[m[j] - '0']++;
        }
    }
    int max = -1;
    for (int i = 0; i < 10; i++) {
        if (max < a[i]) {
            max = a[i];
        }
    }
    cout << max << endl;
    return 0;
}

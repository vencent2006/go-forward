#include <iostream>
using namespace std;
char a[101];
int b[101];

bool isRepeat(int pos) {
    for (int i = 0; i < pos; i++) {
        if (a[i] == a[pos] && a[i + 1] == a[pos + 1]) {
            return true;
        }
    }
    return false;
}

int main() {
    cin >> a;
    int cnt = 0;
    int distinct = 0;
    for (int i = 0; i < strlen(a) - 2; i++) {
        if (a[i] != a[i + 1] && a[i + 1] == a[i + 2]) {
            cout << a[i] << a[i + 1] << a[i + 2] << " ";
            b[i] = 1;
            cnt++;
            if (!isRepeat(i)) {
                distinct++;
            }
        }
    }
    cout << cnt << " " << distinct << endl;
    return 0;
}

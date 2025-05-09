#include <iostream>
using namespace std;

int main() {
    string s;
    cin >> s;
    int cnt = 0;
    int len = s.length();
    for (int i = 0; i + 2 <= len; i++) {
        if (s[i] != s[i + 1] && s[i + 1] == s[i + 2]) {
            // cout << s[i] << s[i + 1] << s[i + 2] << endl;
            cnt++;
        }
    }
    cout << cnt << endl;
    return 0;
}

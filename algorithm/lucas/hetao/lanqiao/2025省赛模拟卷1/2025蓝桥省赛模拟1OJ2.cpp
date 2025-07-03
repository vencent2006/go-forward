#include <iostream>
using namespace std;

int main() {
    string s;
    cin >> s;
    int count = 0;
    for (int i = 0; i <= s.length() - 3; i++) {
        if (s[i] != s[i + 1] && s[i + 1] == s[i + 2]) {
            // cout << s[i] << s[i + 1] << s[i + 2] << " ";
            count++;
        }
    }
    cout << count;
    return 0;
}

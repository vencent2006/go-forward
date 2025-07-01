#include <iostream>
using namespace std;

int main() {
    string s;
    cin >> s;
    int left = 0, right = s.length() - 1;
    while (left < right) {
        char tmp;
        tmp = s[right];
        s[right] = s[left];
        s[left] = tmp;
        left++;
        right--;
    }
    cout << s << endl;
    return 0;
}

#include <iostream>
using namespace std;
// https://htoj.com.cn/cpp/oj/problem/detail?pid=22386398395520&cid=22394268493184&gid=22185991356800
int main() {
    string s;
    cin >> s;
    bool flag = false;
    int left = 0, right = s.length() - 1;
    while (left < right) {
        if (s[left] == s[right]) {
            if (s[left] == '7') {
                flag = true;
            }
            left++;
            right--;
        } else {
            cout << "no";
            return 0;
        }
    }
    if (flag) {
        cout << "yes";
    } else {
        cout << "no";
    }
    return 0;
}

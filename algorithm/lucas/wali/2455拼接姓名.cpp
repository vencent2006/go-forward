#include <iostream>
#include <string>
using namespace std;

int main() {
    string s1, s2;
    cin >> s1 >> s2;
    int l1 = (int)s1.length(), l2 = (int)s2.length();
    bool isLower =false;// s1是否小于s2
    if (l1 < l2) {
        isLower = true;
    }



    for (int i = 0; i < s1.length() && i < s2.length(); i++) {
        if (s1[i] < s2[i]) {
            isLower = true;
            break;
        }
        if (s1[i] > s2[i]) {
            isLower = false;
            break;
        }
    }
    if (isLower) {
        cout << s1 + s2 << endl;
    } else {
        cout << s2 + s1 << endl;
    }
    return 0;
}

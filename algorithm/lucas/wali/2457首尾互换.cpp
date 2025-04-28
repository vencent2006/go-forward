#include <iostream>
using namespace std;
int main() {
    string s;
    cin >> s;
    string s1 = s;
    for (int i = 0; i < s.length(); i++) {
        s1[i] = s[s.length() - i - 1];
    }
    cout<<s+s1<<endl;
    return 0;
}
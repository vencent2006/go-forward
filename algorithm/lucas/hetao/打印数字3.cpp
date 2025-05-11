#include <iostream>
using namespace std;
string zero[5] = {
    "*****",
    "*   *",
    "*   *",
    "*   *",
    "*****"
};
string one[5] = {
    "    *",
    "    *",
    "    *",
    "    *",
    "    *"
};
string two[5] = {
    "*****",
    "    *",
    "*****",
    "*    ",
    "*****"
};
string three[5] = {
    "*****",
    "    *",
    "*****",
    "    *",
    "*****"
};

int main() {
    string s;
    cin >> s;
    int weishu = s.length();
    for (int i = 0; i < 5; i++) {
        for (int j = 0; j < weishu; j++) {
            switch (s[j]) {
                case '0': cout << zero[i];
                    break;
                case '1': cout << one[i];
                    break;
                case '2': cout << two[i];
                    break;
                case '3': cout << three[i];
                    break;
                default: break;
            }
            cout << "   ";
        }
        cout << endl;
    }
    return 0;
}

#include <iostream>
using namespace std;
char zero[5][6] = {
    "*****",
    "*   *",
    "*   *",
    "*   *",
    "*****"
};
char one[5][6] = {
    "    *",
    "    *",
    "    *",
    "    *",
    "    *"
};
char two[5][6] = {
    "*****",
    "    *",
    "*****",
    "*    ",
    "*****"
};
char three[6][6] = {
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
            int idx = s[j] - '0';
            if (idx == 0) {
                cout << zero[i];
            } else if (idx == 1) {
                cout << one[i];
            } else if (idx == 2) {
                cout << two[i];
            } else if (idx == 3) {
                cout << three[i];
            }
            cout << " ";
        }
        cout << endl;
    }
    return 0;
}

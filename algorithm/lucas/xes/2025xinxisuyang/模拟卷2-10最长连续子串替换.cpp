#include <iostream>
using namespace std;

int main() {
    string T, P;
    cin >> T >> P;
    int lenT = T.length();
    int lenP = P.length();
    int minReplace = INT_MAX;

    if (lenP > lenT) {
        cout << -1 << endl;
        return 0;
    }

    for (int i = 0; i < lenT - lenP; i++) {
        int replace = 0;
        bool match = true;
        for (int j = 0; j < lenP; j++) {
            if (T[i + j] == '?') {
                replace++;
            } else if (T[i + j] != P[j]) {
                match = false;
                break;
            }
        }
        if (match) {
            if (replace < minReplace) {
                minReplace = replace;
            }
        }
    }

    if (minReplace == INT_MAX) {
        cout << -1 << endl;
    } else {
        cout << minReplace << endl;
    }
    return 0;
}

#include <iostream>
using namespace std;

int main() {
    char s;
    cin >> s;
    int add = 0;


    switch (s) {
        case 'A':
            add = 500;
            break;
        case 'B':
            add = 300;
            break;
        case 'C':
            add = 0;
            break;
        case 'D':
            add = -200;
            break;
        case 'E':
            add = -500;
            break;
        default:
            add = 0;
            cout << "error" << endl;
            break;
    }


    cout << 8000 + add << endl;

    return 0;
}

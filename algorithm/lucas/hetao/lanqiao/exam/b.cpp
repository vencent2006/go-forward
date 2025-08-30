#include <iostream>

using namespace std;

int main() {
    int i = 1, t = 0;
    while (i * i < 30) {
        t += i;
        i += 2;
    }
    cout << t << endl;
    cout << (char) ('F' + 4);
    return 0;
}

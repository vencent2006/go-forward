#include <iostream>

using namespace std;

int seconds(int h, int m, int s, char wu) {
    if (wu == 'P') {
        h += 12;
    }
    return h * 3600 + m * 60 + s;
}

int main() {
    int h, m, s;
    char wu;
    cin >> h >> m >> s >> wu;
    cout << seconds(h, m, s, wu) << endl;
    return 0;
}

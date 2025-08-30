#include <iostream>

using namespace std;

int count(int area) {
    int cnt = 0;
    for (int i = 1; i * i <= area; i++) {
        if (area % i == 0) {
            cnt++;
        }
    }
    return cnt;
}

int main() {
    int area;
    cin >> area;
    cout << count(area) << endl;
    return 0;
}

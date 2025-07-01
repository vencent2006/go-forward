#include <iostream>
using namespace std;

bool isA(int n) {
    return n % 3 == 0;
}

bool isB(int n) {
    return n % 5 == 0 && n % 3 != 0;
}

bool isC(int n) {
    return n % 5 != 0 && n % 3 != 0;
}

int main() {
    int n;
    cin >> n;
    int sumA = 0, sumB = 0, sumC = 0;
    for (int i = 1; i <= n; i++) {
        if (isA(i)) {
            sumA += i;
        } else if (isB(i)) {
            sumB += i;
        } else if (isC(i)) {
            sumC += i;
        }
    }
    cout << sumA << " " << sumB << " " << sumC << endl;

    return 0;
}

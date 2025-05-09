#include <iostream>
using namespace std;

int shuweihe(int n) {
    int sum = 0;
    while (n != 0) {
        int m = n % 10;
        sum += m;
        n = n / 10;
    }
    return sum;
}

int main() {
    int n;
    cin >> n;
    int max = -1;
    int sum = shuweihe(n);
    if (max < sum) {
        max = sum;
    }
    cout << sum << endl;
    return 0;
}

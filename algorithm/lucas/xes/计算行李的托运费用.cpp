#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;
    float ans = 0.0;
    if (n <= 30) {
        ans = n * 0.2;
    } else if (n <= 60) {
        ans = 30 * 0.2 + (n - 30) * 0.6;
    } else {
        cout << "NO";
        return 0;
    }
    printf("%.2f", ans);
    return 0;
}

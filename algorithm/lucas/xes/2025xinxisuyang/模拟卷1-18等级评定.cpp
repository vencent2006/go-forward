#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;
    for (int i = 1; i <= n; i++) {
        int score;
        cin >> score;
        if (score >= 95) {
            cout << "优" << " ";
        } else if (score >= 85) {
            cout << "良" << " ";
        } else if (score >= 75) {
            cout << "中" << " ";
        }
    }
    return 0;
}

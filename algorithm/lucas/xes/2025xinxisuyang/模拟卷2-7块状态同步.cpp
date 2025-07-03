#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;
    string A, B;
    cin >> A >> B;
    int count = 0;
    bool isFlip = false; // 记录是否翻转过
    for (int i = 0; i < n; i++) {
        if (A[i] != B[i]) {
            if (!isFlip) {
                isFlip = true;
                count++;
            }
        } else {
            isFlip = false;
        }
    }
    cout << count << endl;
    return 0;
}

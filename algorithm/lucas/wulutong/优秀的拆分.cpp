#include <iostream>
using namespace std;

int main() {
  int n;
  cin >> n;
  if (n & 1) {
    // 奇数
    cout << -1;
    return 0;
  }

  for (int i = 31; i >= 1; i--) {
    int t = (1 << i);
    if (t & n) {
      // 对应的位是1
      cout << t << " ";
    }
  }
  return 0;
}

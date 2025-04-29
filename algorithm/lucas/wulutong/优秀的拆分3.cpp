#include <iostream>
using namespace std;


int main() {
  int n;
  if (freopen("power.in", "r", stdin) == NULL) {
    // 从power.in读入数据
    cerr << "Failed to open input file." << endl;
    return 1;
  }
  freopen("power.out", "w", stdout); // 输出到power.out
  cin >> n;
  if (n & 1) {
    // 奇数
    cout << -1;
    return 0;
  }

  for (int i = 31; i >= 1; i--) {
    int t = (1 << i);
    if (t & n) {
      // 对应的位是1, 从大到小遍历2的整数次方，存在就输出
      cout << t << " ";
    }
  }
  return 0;
}

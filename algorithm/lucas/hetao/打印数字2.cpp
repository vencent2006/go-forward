#include <iostream>
#include <vector>
#include <string>
using namespace std;

int main() {
    // 初始化数字的5x5表示
    vector<vector<string> > digits(4);
    // 初始化数字0
    digits[0].push_back("*****");
    digits[0].push_back("*   *");
    digits[0].push_back("*   *");
    digits[0].push_back("*   *");
    digits[0].push_back("*****");

    // 初始化数字1
    digits[1].push_back("    *");
    digits[1].push_back("    *");
    digits[1].push_back("    *");
    digits[1].push_back("    *");
    digits[1].push_back("    *");

    // 初始化数字2
    digits[2].push_back("*****");
    digits[2].push_back("    *");
    digits[2].push_back("*****");
    digits[2].push_back("*    ");
    digits[2].push_back("*****");

    // 初始化数字3
    digits[3].push_back("*****");
    digits[3].push_back("    *");
    digits[3].push_back("*****");
    digits[3].push_back("    *");
    digits[3].push_back("*****");

    string n;
    cin >> n;


    // 检查输入是否在有效范围
    for (int i = 0; i < sizeof(n); i++) {
        int t = n[i] - '0';
        for (int j = 0; j < 5; ++j) {
            cout << digits[t][j] << endl;
        }
    }

    return 0;
}

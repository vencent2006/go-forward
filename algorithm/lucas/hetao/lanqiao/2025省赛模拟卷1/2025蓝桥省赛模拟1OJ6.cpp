#include <iostream>
using namespace std;

// 自定义比较函数，用于降序排序
bool cmp(char a, char b) {
    return a > b; // 如果a大于b，返回true，这样a会排在b前面
}

int main() {
    string s;
    cin >> s;


    int len = s.length();
    int i = len - 2;
    // 从右向左找第一个下降点（即s[i] > s[i+1]的位置）
    while (i >= 0 && s[i] <= s[i + 1]) {
        i--;
    }

    // 如果没找到下降点，说明没有更小的排列
    if (i < 0) {
        cout << -1 << endl;
        return 0;
    }

    // 在i位置右侧找比s[i]小的最大数字
    char maxDigit = '0';
    int pos = -1;
    for (int j = i + 1; j < len; j++) {
        if (s[j] < s[i] && s[j] >= maxDigit) {
            maxDigit = s[j];
            pos = j;
        }
    }

    // 交换位置i和pos的数字
    swap(s[i], s[pos]);
    // 对i位置右侧进行降序排序（使用自定义的cmp函数）
    sort(s.begin() + i + 1, s.end(), cmp);

    // 去除前导零
    int start = 0;
    while (start < s.length() && s[start] == '0') {
        start++;
    }

    // 检查结果是否有效
    // 所有位都是0，无效结果
    // 输出去除前导零后的结果
    if (start == s.length()) {
        cout << -1 << endl;
    } else {
        for (int j = start; j < s.length(); j++) {
            cout << s[j];
        }
    }
    return 0;
}

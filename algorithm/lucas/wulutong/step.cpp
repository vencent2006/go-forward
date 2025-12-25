#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

int minStepsToSameString(const string &s) {
    if (s.empty()) return 0;

    // 统计26个小写字母的出现频率
    vector<int> freq(26, 0);

    for (char c: s) {
        int index = c - 'a';  // 将字符映射到0-25
        freq[index]++;
    }

    // 找出最大频率
    int maxCount = 0;
    for (int count: freq) {
        if (count > maxCount) {
            maxCount = count;
        }
    }

    // 最少步骤：将其他所有字符改为出现最多的字符
    // 总共有 s.length() - maxCount 个字符需要改变
    return s.length() - maxCount;
}

int main() {
    string input;

    cout << "输入一串小写字符串: ";
    cin >> input;

    // 验证输入
    for (char c: input) {
        if (c < 'a' || c > 'z') {
            cout << "错误：输入必须只包含小写字母！" << endl;
            return 1;
        }
    }

    int steps = minStepsToSameString(input);
    cout << "最少需要 " << steps << " 步" << endl;

    // 解释计算过程
    cout << "\n计算过程：" << endl;
    cout << "字符串长度: " << input.length() << endl;

    // 统计并显示频率
    vector<int> freq(26, 0);
    for (char c: input) {
        freq[c - 'a']++;
    }

    char mostFrequentChar = 'a';
    int maxFreq = 0;
    for (int i = 0; i < 26; i++) {
        if (freq[i] > maxFreq) {
            maxFreq = freq[i];
            mostFrequentChar = 'a' + i;
        }
    }

    cout << "最频繁的字符: '" << mostFrequentChar << "' (出现 " << maxFreq << " 次)" << endl;
    cout << "需要改变的字符数: " << input.length() << " - " << maxFreq << " = " << steps << endl;

    return 0;
}
#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;
    int cishu[101] = {0}; // 0到100,都初始化为0
    int max = -1;
    int min = 101;
    int sum = 0;

    // 输入：遍历输入的数
    for (int i = 1; i <= n; i++) {
        int score;
        cin >> score;
        if (score > max) { max = score; } // 是最大值吗
        if (score < min) { min = score; } // 是最小值吗
        cishu[score]++; // 对应数的次数加1
        sum += score; // 总和加1
    }

    // 输出1：出现过的分数和次数
    for (int i = 0; i <= 100; i++) {
        if (cishu[i] > 0) {
            cout << i << ":" << cishu[i] << endl;
        }
    }
    // 输出2：最高分 最低分 平均分
    printf("最高分:%d 最低分:%d 平均分:%.2f", max, min, (double) sum / (double) n);

    return 0;
}

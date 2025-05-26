#include <iostream>
using namespace std;

int findDigitInFraction(int a, int b, int x) {
    if (b == 0) {
        return -1; // 除数为0，无效
    }
    // 处理整数部分
    int integerPart = a / b;
    a %= b;
    // 处理小数部分
    for (int i = 1; i <= x; ++i) {
        a *= 10;
        int digit = a / b;
        if (i == x) {
            return digit;
        }
        a %= b;
        if (a == 0) {
            return 0; // 如果小数部分已经结束，后续位数都是0
        }
    }
    return -1; // 不应该执行到这里
}

int main() {
    int a, b, x;
    cout << "请输入三个整数a, b, x: ";
    cin >> a >> b >> x;
    if (b == 0) {
        cout << "错误：除数不能为0" << endl;
        return 1;
    }
    if (x <= 0) {
        cout << "错误：x必须是正整数" << endl;
        return 1;
    }

    int digit = findDigitInFraction(a, b, x);
    cout << a << "/" << b << " 的小数表示中第 " << x << " 位数字是: " << digit << endl;

    return 0;
}

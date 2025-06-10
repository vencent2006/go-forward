#include <iostream>
using namespace std;

// 打印整数的二进制表示，可指定位数
void printBinary(int num, int bitsCount = 32) {
    for (int i = bitsCount - 1; i >= 0; i--) {
        cout << ((num >> i) & 1);
    }
}

// 计算并返回原码、反码、补码
struct NumberRepresentations {
    int original;
    int inverse;
    int complement;
};

NumberRepresentations getRepresentations(int num) {
    NumberRepresentations result;
    result.original = num;
    if (num >= 0) {
        result.inverse = num;
        result.complement = num;
    } else {
        result.inverse = (1 << 31) | (~(-num) & 0x7FFFFFFF);
        result.complement = num; // 在计算机中，负数补码就是其存储形式
    }
    return result;
}

// 打印原码、反码和补码，可指定位数
void printOriginalInverseComplement(int num, int bitsCount = 32) {
    NumberRepresentations reps = getRepresentations(num);
    cout << "原码: ";
    printBinary(reps.original, bitsCount);
    cout << "\n反码: ";
    printBinary(reps.inverse, bitsCount);
    cout << "\n补码: ";
    printBinary(reps.complement, bitsCount);
    cout << endl;
}

int main() {
    const int a[6] = {-27, 35, 44, -15, -18, -23};
    for (int i = 0; i < 6; i++) {
        cout << i << ": start with " << a[i] << endl;
        printOriginalInverseComplement(a[i], 8); // 输出 8 位表示
    }
    return 0;
}

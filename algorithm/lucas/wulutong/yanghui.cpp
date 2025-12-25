#include <iostream>
#include <vector>
using namespace std;

vector<vector<int>> generatePascalTriangle(int n) {
    vector<vector<int>> triangle;

    if (n <= 0) return triangle;

    // 第一行
    triangle.push_back({1});

    // 生成后续行
    for (int i = 1; i < n; i++) {
        vector<int> row(i + 1, 1);  // 每行首尾都是1

        // 计算中间的值：row[j] = 上一行的row[j-1] + 上一行的row[j]
        for (int j = 1; j < i; j++) {
            row[j] = triangle[i-1][j-1] + triangle[i-1][j];
        }

        triangle.push_back(row);
    }

    return triangle;
}

void printPascalTriangle(const vector<vector<int>>& triangle) {
    for (const auto& row : triangle) {
        for (int num : row) {
            cout << num << " ";
        }
        cout << endl;
    }
}

int main() {
    int n;
    cout << "请输入要生成的杨辉三角形的行数: ";
    cin >> n;

    vector<vector<int>> triangle = generatePascalTriangle(n);
    cout << "杨辉三角形前 " << n << " 行:" << endl;
    printPascalTriangle(triangle);

    return 0;
}
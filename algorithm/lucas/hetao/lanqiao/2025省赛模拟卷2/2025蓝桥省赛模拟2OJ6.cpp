#include <iostream>
using namespace std;
int m, n;
int a[1005][1005];
//返回当前岛屿面积，标记已经访问的地方
int dfs(int i, int j) {
    if (i < 1 || i > m || j < 1 || j > n || a[i][j] == 0) return 0;

    a[i][j] = 0;
    return 1 + dfs(i - 1, j) + dfs(i + 1, j) + dfs(i, j - 1) + dfs(i, j + 1);
}

int main() {
    cin >> m >> n;
    //输入地形图
    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            cin >> a[i][j];
        }
    }
    int maxx = -9999, minn = 9999;
    //计算所有岛屿的面积
    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            if (a[i][j] == 1) {
                int area = dfs(i, j);
                maxx = max(maxx, area);
                minn = min(minn, area);
            }
        }
    }
    //处理没有岛屿的情况
    if (minn == 9999) minn = 0;
    cout << maxx - minn;
    return 0;
}

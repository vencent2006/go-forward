#include <iostream>
#include <algorithm>
using namespace std;

struct stu {
    int score;
    string name;
};

stu a[100];

bool cmp(stu x, stu y) {
    return x.score > y.score;
}

int main() {
    int n;
    cin >> n;
    for (int i = 0; i < n; i++) {
        cin >> a[i].score >> a[i].name;
    }
    sort(a, a + n, cmp);
    cout << a[0].name;

    return 0;
}

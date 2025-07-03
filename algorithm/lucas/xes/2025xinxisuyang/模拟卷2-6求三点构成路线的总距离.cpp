#include <iostream>
#include <cmath>
using namespace std;

typedef struct {
    float x;
    float y;
} Node;

int main() {
    Node arr[3];
    for (int i = 0; i < 3; i++) {
        cin >> arr[i].x >> arr[i].y;
    }
    float distance = sqrt((arr[0].x - arr[1].x) * (arr[0].x - arr[1].x) + (arr[0].y - arr[1].y) * (arr[0].y - arr[1].y))
                     +
                     sqrt((arr[1].x - arr[2].x) * (arr[1].x - arr[2].x) + (arr[1].y - arr[2].y) * (
                              arr[1].y - arr[2].y));
    printf("%.2f", distance);
    return 0;
}

#include <iostream>
using namespace std;
int getSum(int num){
    int sum = 0;
    while(num>0){
        sum += num % 10;
        num /= 10;
    }
    return sum;
}
int main(){
    int n;
    cin >> n;
    int max = -1, sum = 0;
    for (int i = 0; i < n; i++) {
        int num;
        cin >> num;
        sum = getSum(num);
        if (sum > max) {
          max = sum;
        }
    }
    cout << max << endl;
    return 0;
}

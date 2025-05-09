#include <iostream>
using namespace std;

int main() {
    int startHour, startMinute, endHour, endMinute;
    cin >> startHour >> startMinute >> endHour >> endMinute;
    int start = startHour * 60 + startMinute;
    int end = endHour * 60 + endMinute;
    cout << end - start;
    return 0;
}

#include <iostream>

int multiply(int x, int y);

int main() {
    int a = 4;
    int b = 5;
    int result = multiply(a, b);
    std::cout << "The result is: " << result << std::endl;
    return 0;
}

int multiply(int x, int y) {
    return x * y;
}

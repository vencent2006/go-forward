# This is a sample Python script.

# Press ⌃R to execute it or replace it with your code.
# Press Double ⇧ to search everywhere for classes, files, tool windows, actions, and settings.
import time
import turtle


def print_hi(name):
    # Use a breakpoint in the code line below to debug your script.
    print(f'Hi, {name}')  # Press ⌘F8 to toggle the breakpoint.


def qiuhe(a, b):
    return a + b


# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    turtle.color("red", "yellow")
    turtle.begin_fill()
    for _ in range(100):
        turtle.forward(200)
        turtle.left(170)
    turtle.end_fill()

    turtle.mainloop()
# See PyCharm help at https://www.jetbrains.com/help/pycharm/

import sys

filename_a = sys.argv[1]
filename_b = sys.argv[2]


def loadFileInSet(filename: str):
    set_local = set()
    with open(filename, "r") as file:
        lines = file.readlines()
        # print(f'{filename} 行数 =', len(lines))
        for line in lines:
            line = line.strip()  # 去除每行的换行符
            # 进行需要的操作
            set_local.add(line)
    return set_local


def printSet(s: set):
    for x in s:
        print(x)


set_a = loadFileInSet(filename_a)
set_b = loadFileInSet(filename_b)
# print(len(set_a))
# print(len(set_b))
set_diff = set_a.difference(set_b)
# print(len(set_diff))
printSet(set_diff)

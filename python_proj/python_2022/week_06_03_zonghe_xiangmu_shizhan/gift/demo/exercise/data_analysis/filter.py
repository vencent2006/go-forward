payed_map = {}

with open("payed.txt", "r") as file:
    lines = file.readlines()
    print('payed.txt 行数 =', len(lines))
    for line in lines:
        line = line.strip()  # 去除每行的换行符
        # 进行需要的操作
        payed_map[line] = 1
        # print(line, end="\n")

print('支付了的订单数 =', len(payed_map), end="\n\n")


def filterPayed(filter_map, list_filename):
    local_map = {}
    local_set = set()
    with open(list_filename, "r") as file:
        lines = file.readlines()
        print(f'{list_filename} 行数 =', len(lines))
        for line in lines:
            line = line.strip()  # 去除每行的换行符
            # 进行需要的操作
            local_set.add(line)
            if line in filter_map:
                local_map[line] = 1
    print(f'{list_filename} 去重 =', len(local_set))
    print(f'{list_filename} 匹配 =', len(local_map), end="\n\n")


filterPayed(payed_map, "xiquezhuangban.txt")
filterPayed(payed_map, "fulishang.txt")

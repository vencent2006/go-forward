list = []
for i in range(5):
    list.append(int(input()))
for i in list:
    if i == 6 or i == 8:
        print('幸运数字')
    else:
        print('普通数字')

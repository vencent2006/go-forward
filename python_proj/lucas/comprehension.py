if __name__ == '__main__':
    # 长度大于3的字符串要大写
    names = ['vincent', 'lucas', 'lee', 'jessie']
    namesUpper = [name.upper() for name in names if len(name) > 3]
    print(namesUpper)

    # 能被3整除的数
    multiple = [i for i in range(30) if i % 3 == 0]
    print(multiple)

    # 字典推导式
    dic = {i: i ** 2 for i in (2, 4, 6)}
    print(dic)

    listDemo = ['Google', 'Runoob', 'Taobao', 'Vincent']
    newDict = {l: len(l) for l in listDemo}
    print(newDict)

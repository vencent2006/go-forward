# 传不可变的参数
def change(a):
    print(id(a))
    a = 10
    print(id(a))  # 与上一个 id(a) 的值不一样


a = 1
print(id(a))
change(a)


# 传 可变的参数
def changeme(mylist):
    """修改传入的列表"""
    mylist.append([1, 2, 3, 4])
    print("函数内 取值: ", mylist)
    return


mylist = [10, 20, 30]
changeme(mylist)
print("函数外 取值: ", mylist)

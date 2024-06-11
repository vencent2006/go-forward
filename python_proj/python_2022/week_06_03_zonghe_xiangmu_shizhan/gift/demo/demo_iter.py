# coding:utf-8
def flatten(nested):
    try:
        for sublist in nested:
            for element in flatten(sublist):
                yield element
    except TypeError:
        yield nested


def flatten2(nested):
    try:
        # 不迭代类似于字符串的对象:
        try:
            nested + ''
        except TypeError:
            pass
        else:
            raise TypeError
        # 继续迭代
        for sublist in nested:
            for element in flatten2(sublist):
                yield element
    except TypeError:
        yield nested


if __name__ == '__main__':
    print(list(flatten([[[1], 2], 3, 4, [5, [6, 7]], 8])))
    print(list(flatten2([[[1], 2], 3, 4, [5, [6, 7]], 8])))
    print(list(flatten2(['foo', ['bar', ['baz']]])))

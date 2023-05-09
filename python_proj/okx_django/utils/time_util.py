import time


def get_current_second(to_str=False):
    """
    获取当前时间戳秒
    print (t)                       #原始时间数据
    print (int(t))                  #秒级时间戳
    print (int(round(t * 1000)))    #毫秒级时间戳
    print (int(round(t * 1000000))) #微秒级时间戳
    :return:

    """
    t = time.time()
    if to_str:
        return str(int(t))
    else:
        return int(t)


if __name__ == '__main__':
    print(get_current_second(True))
    print(get_current_second(False))

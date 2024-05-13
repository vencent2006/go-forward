def echo_key(key: str, info: map):
    if key in info.keys():
        echo(info[key])
    pass


def insert_db(info: dict):
    echo(info['uid'])
    echo(",")
    echo_key('1', info)
    echo(",")
    echo_key('3', info)
    echo(",")
    echo_key('6', info)
    echo(",")
    echo_key('12', info)
    print()
    pass


def echo(i: str):
    print(i, end="")


old = ''
temp_dict = {}
print("vuid,月费,3月,6月,12个月")
with open("dingyue_raw.txt", "r") as file:
    i = -1
    for line in file:
        i = i + 1
        # if i > 100:
        #     break

        record = line.strip().split()
        # print(record)
        if i == 0:
            old = record[0]
            temp_dict['uid'] = record[0]

        if old == record[0]:
            # 还是当前的记录
            temp_dict[record[1]] = record[2]
            pass
        else:
            # 新纪录
            if len(temp_dict) > 0:
                # 插入到数据库
                insert_db(temp_dict)
            # reset
            old = record[0]
            temp_dict = {'uid': record[0], record[1]: record[2]}

            pass

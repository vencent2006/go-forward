# coding:utf-8

"""
    2-1
    1. 导入 user.json 文件检查
    2. 导入 gift.json 文件检查

    3-1
    1. 确定用户表中每个用户的信息字段
    2. 读取 user_json 文件
    3. 写入 user_json 文件 （检测该用户是否存在）， 存在则不可写入

    username 姓名
    role normal or admin
    active True or False
    create_time timestamp
    update_time timestamp
    gifts []

    key:value
    username:{username, role, active}

    Python中的*（星号）和**(双星号）完全详解 https://blog.csdn.net/m0_48891301/article/details/134004143
"""
import json
import os.path
import time

from common.error import UserExistsError
from common.utils import check_file, timestamp_to_str


class Base(object):
    def __init__(self, user_json, gift_json):
        self.user_json = user_json
        self.gift_json = gift_json

        self.__check_user_json()
        self.__check_gift_json()
        return

    def __check_user_json(self):
        check_file(self.user_json)
        return

    def __check_gift_json(self):
        check_file(self.gift_json)
        return

    def __read_users(self, time_to_str=False):
        with open(self.user_json, 'r') as f:
            data = json.loads(f.read())

        if time_to_str == True:
            for username, v in data.items():
                v['create_time'] = timestamp_to_str(v['create_time'])
                v['update_time'] = timestamp_to_str(v['update_time'])
                data[username] = v
        return data

    def write_user(self, **user):
        if 'username' not in user:
            raise ValueError('missing username')
        if 'role' not in user:
            raise ValueError('missing role')

        user['active'] = True
        user['create_time'] = time.time()
        user['update_time'] = time.time()
        user['gifts'] = []

        users = self.__read_users()
        print(users)
        # return
        if user['username'] in users:
            raise UserExistsError(message='username %s already existed' % user['username'])
        users.update(
            {user['username']: user}
        )
        json_users = json.dumps(users)
        with open(self.user_json, 'w') as f:
            f.write(json_users)
        return


if __name__ == '__main__':
    gift_path = os.path.join(os.getcwd(), 'storage', 'gift.json')
    user_path = os.path.join(os.getcwd(), 'storage', 'user.json')

    print(gift_path)
    print(user_path)

    base = Base(user_json=user_path, gift_json=gift_path)
    base.write_user(username='dewei', role='admin')

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
    Python | 一文简单搞懂json.dump()与json.dumps()的区别 https://blog.csdn.net/weixin_46264660/article/details/130238426

    4-1
    1. role的修改
    2. active的修改
    3. delete_user

    5-1
    1. gifts 奖品结构的确定
    2. gifts 奖品的读取
    3. gifts 添加
    4. gifts 初始化

    奖池的存储格式
    {
        level1: {
            level1: {
                gift_name1: {
                    name: xx
                    count: xx
                },
                gift_name2: {
                    name: xx
                    count: xx
                }
            },
            level2: {},
            level3: {}
        },
        level2: {
            level1: {},
            level2: {},
            level3: {}
        }
    }
"""
import json
import os.path
import time

from common.consts import ROLES, FIRSTLEVELS, SECONDLEVELS
from common.error import UserExistsError, RoleError, LevelError, GiftCountError
from common.utils import check_file, timestamp_to_str


class Base(object):
    def __init__(self, user_json, gift_json):
        self.user_json = user_json
        self.gift_json = gift_json

        # 检查数据
        self.__check_user_json()
        self.__check_gift_json()

        # 初始化
        self.__init_gifts()
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

    def __write_user(self, **user):
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

    def __change_role(self, username, role):
        users = self.__read_users()
        user = users[username]  # {'username':{role, create_time}}
        if not user:
            return False
        if role not in ROLES:
            raise RoleError('not use role %s' % role)
        user['role'] = role
        user['update_time'] = time.time()
        users['username'] = user

        json_data = json.dumps(users)
        with open(self.user_json, 'w') as f:
            f.write(json_data)
        return True

    def __change_active(self, username):
        users = self.__read_users()
        user = users.get(username)
        if not user:
            return False
        user['active'] = not user['active']
        user['update_time'] = time.time()
        users['username'] = user

        json_data = json.dumps(users)
        with open(self.user_json, 'w') as f:
            f.write(json_data)
        return True

    def __delete_user(self, username):
        users = self.__read_users()
        user = users.get(username)
        if not user:
            return False
        delete_user = users.pop(username)
        json_data = json.dumps(users)
        with open(self.user_json, 'w') as f:
            f.write(json_data)
        return delete_user

    def read_gift(self):
        with open(self.gift_json, 'r') as f:
            data = json.loads(f.read())
        return data

    def __init_gifts(self):
        data = {
            'level1': {
                'level1': {},
                'level2': {},
                'level3': {}
            },
            'level2': {
                'level1': {},
                'level2': {},
                'level3': {}
            },
            'level3': {
                'level1': {},
                'level2': {},
                'level3': {}
            },
            'level4': {
                'level1': {},
                'level2': {},
                'level3': {}
            }
        }
        gifts = self.read_gift()
        if len(gifts) != 0:
            return
        json_data = json.dumps(data)
        with open(self.gift_json, 'w') as f:
            f.write(json_data)
        return

    def __write_gift(self, first_level, second_level, gift_name, gift_count):
        if first_level not in FIRSTLEVELS:
            raise LevelError('first level %s not exist' % first_level)
        if second_level not in SECONDLEVELS:
            raise LevelError('second level %s not exist' % second_level)
        if gift_count <= 0:
            raise GiftCountError('gift count %s invalid ' % gift_count)
        gifts = self.read_gift()
        current_gift_pool = gifts[first_level]
        current_second_gift_pool = current_gift_pool[second_level]
        if gift_name in current_second_gift_pool:
            # 如果存在就累加值
            current_second_gift_pool[gift_name]['count'] = current_second_gift_pool[gift_name]['count'] + gift_count
        else:
            # 如果不存在，就设置值
            current_second_gift_pool[gift_name] = {'name': gift_name, 'count': gift_count}

        # 写回数据
        gifts[first_level] = current_second_gift_pool
        json_data = json.dumps(gifts)
        with open(self.gift_json, 'w') as f:
            f.write(json_data)
        return


if __name__ == '__main__':
    gift_path = os.path.join(os.getcwd(), 'storage', 'gift.json')
    user_path = os.path.join(os.getcwd(), 'storage', 'user.json')

    print(gift_path)
    print(user_path)

    base = Base(user_json=user_path, gift_json=gift_path)
    # base.write_user(username='dewei', role='admin')
    # result = base.change_role(username='dewei', role='admin')
    # result = base.delete_user(username='dewei')
    result = base.read_gift()
    print(result)

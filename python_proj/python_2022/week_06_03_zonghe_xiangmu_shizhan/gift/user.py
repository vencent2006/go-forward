# coding: utf-8

"""
    9-1
    1: user 类的初始化
    2: get_user （时间的转变）
    3：查看奖品列表
"""
import os

from base import Base
from common.consts import ROLE_NORMAL
from common.error import NotUserError, RoleError, UserActiveError
from common.utils import timestamp_to_str


class User(Base):
    def __init__(self, username, user_json, gift_json):
        self.username = username
        super().__init__(user_json=user_json, gift_json=gift_json)
        self.get_user()

    def get_user(self):
        users = self._Base__read_users()
        if self.username not in users:
            raise NotUserError('not user %s' % self.username)
        current_user = users.get(self.username)

        if current_user.get('active') == False:
            raise UserActiveError('the user %s not active' % self.username)
        if current_user.get('role') != ROLE_NORMAL:
            raise RoleError('permission by normal')

        self.user = current_user
        self.role = current_user.get('role')
        self.name = current_user.get('username')
        self.gifts = current_user.get('gifts')
        self.create_time = timestamp_to_str(current_user.get('create_time'))

    def get_gifts(self):
        gifts = self._Base__read_gifts()
        # print(gifts)
        gift_list = []
        # print('********')
        for level_one, level_one_pool in gifts.items():
            # print(level_one_pool)
            for level_two, level_two_pool in level_one_pool.items():
                # print(level_two, level_two_pool)
                for gift_name, gift_info in level_two_pool.items():
                    # print(gift_info)
                    # gift_info.pop('count')
                    gift_list.append(gift_info.get('name'))
        return gift_list


if __name__ == '__main__':
    gift_path = os.path.join(os.getcwd(), 'storage', 'gift.json')
    user_path = os.path.join(os.getcwd(), 'storage', 'user.json')
    user = User(username='lucas', user_json=user_path, gift_json=gift_path)
    # print(user.name, user.role, user.create_time, user.gifts)
    gift_list = user.get_gifts()
    print(gift_list)

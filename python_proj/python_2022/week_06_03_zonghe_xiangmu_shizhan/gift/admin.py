# coding:utf-8


"""
    7-1
    1:admin类的搭建
    2:获取用户函数(包含获取身份)
    3:添加用户(判断当前身份是否是管理员)
    4:冻结与恢复用户
    5:修改用户身份

    8-1
    1: admin的验证（只有admin的用户才能用这个类）
    2：任何函数都应该动态的更新getuser
    3：奖品的添加
    4：奖品的删除
    5：奖品数量的更新（同步base调整）
"""
import os

from base import Base
from common.consts import ROLE_ADMIN
from common.error import NotUserError, UserActiveError, RoleError


class Admin(Base):
    def __init__(self, username, user_json, gift_json):
        self.username = username
        super().__init__(user_json, gift_json)
        self.get_user()

    def get_user(self):
        users = self._Base__read_users()  # 还是有办法获得父类的私有方法
        print(users)
        current_user = users.get(self.username)
        if current_user == None:
            raise NotUserError('not user %s' % self.username)
        if current_user.get('active') == False:
            raise UserActiveError('the user %s not active' % self.username)
        if current_user.get('role') != ROLE_ADMIN:
            raise RoleError('permission by admin')

        self.user = current_user
        self.role = current_user.get('role')
        self.name = current_user.get('username')
        self.active = current_user.get('active')

    def add_user(self, username, role):
        self.__check_admin()
        self._Base__write_user(username=username, role=role)

    def update_user_active(self, username):
        self.__check_admin()
        # 正常应该不能修改自己的active
        if username == self.username:
            raise Exception('can not update yourself')
        self._Base__change_active(username=username)

    def update_user_role(self, username, role):
        self.__check_admin()
        # 正常应该不能修改自己的active
        if username == self.username:
            raise Exception('can not update yourself')
        self._Base__change_role(username=username, role=role)

    def add_gift(self, first_level, second_level, gift_name, gift_count):
        self.__check_admin()
        self._Base__write_gift(first_level=first_level, second_level=second_level, gift_name=gift_name,
                               gift_count=gift_count)

    def delete_gift(self, first_level, second_level, gift_name):
        self.__check_admin()
        self._Base__delete_gift(first_level=first_level, second_level=second_level, gift_name=gift_name)

    def update_gift(self, first_level, second_level, gift_name, gift_count):
        self.__check_admin()
        self._Base__gift_update(first_level=first_level, second_level=second_level, gift_name=gift_name,
                                gift_count=gift_count, is_admin=True)

    def __check_admin(self):
        self.get_user()


if __name__ == '__main__':
    gift_path = os.path.join(os.getcwd(), 'storage', 'gift.json')
    user_path = os.path.join(os.getcwd(), 'storage', 'user.json')

    print(gift_path)
    print(user_path)

    admin = Admin('dewei', user_json=user_path, gift_json=gift_path)
    # print(admin.name)
    # admin.add_user(username='dewei', role='admin')
    # admin.add_user(username='lucas', role='normal')
    # admin.update_user_active('lucas')
    # admin.update_user_role('lucas', ROLE_NORMAL)
    admin.add_gift(first_level='level1', second_level='level1', gift_name='糖豆', gift_count=1000)
    # admin.delete_gift(first_level='level1', second_level='level1', gift_name='糖豆')
    admin.update_gift(first_level='level1', second_level='level1', gift_name='糖豆', gift_count=1)

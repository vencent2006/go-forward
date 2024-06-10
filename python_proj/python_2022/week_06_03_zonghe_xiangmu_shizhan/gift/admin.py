# coding:utf-8


"""
    7-1
    1:admin类的搭建
    2:获取用户函数(包含获取身份)
    3:添加用户(判断当前身份是否是管理员)
    4:冻结与恢复用户
    5:修改用户身份
"""
import os

from base import Base
from common.consts import ROLE_ADMIN, ROLE_NORMAL
from common.error import NotUserError, UserActiveError


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

    def __check_admin(self):
        if self.role != ROLE_ADMIN:
            raise Exception('permission')


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
    admin.update_user_role('lucas', ROLE_NORMAL)

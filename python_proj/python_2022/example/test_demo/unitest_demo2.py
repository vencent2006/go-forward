import unittest


def login(username, password):
    if username == 'kira' and password == '123':
        res = {"code": 200, "msg": "登录成功"}
        return res
    return {"code": 400, "msg": "登录失败"}


class TestLogin(unittest.TestCase):

    def test_login_success(self):
        """测试登录成功"""
        print("enter test_login_success")
        test_data = {"username": "kira", "password": "test"}
        expect_data = {"code": 200, "msg": "登录成功"}
        res = login(**test_data)
        self.assertEqual(res, expect_data)

    def test_login_error_with_error_password(self):
        """账号正确，密码错误，登录失败"""
        print("enter test_login_error_with_error_password")
        test_data = {"username": "kira", "password": "12345"}
        expect_data = {"code": 400, "msg": "登录失败"}
        res = login(**test_data)
        self.assertEqual(res, expect_data)

    # 更多测试函数类似...


if __name__ == '__main__':
    unittest.main()

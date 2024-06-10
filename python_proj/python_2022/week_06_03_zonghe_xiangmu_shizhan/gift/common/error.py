# coding:utf-8

# ----- file error -----
class NotPathError(Exception):
    def __init__(self, message):
        self.message = message


class FormatError(Exception):
    def __init__(self, message='need json format'):
        self.message = message


class NotFileError(Exception):
    def __init__(self, message='this is not a file'):
        self.message = message


# ----- user error -----
class UserExistsError(Exception):
    def __init__(self, message):
        self.message = message


class RoleError(Exception):
    def __init__(self, message):
        self.message = message


# ----- gift error -----

class LevelError(Exception):
    def __init__(self, message):
        self.message = message


class GiftCountError(Exception):
    def __init__(self, message):
        self.message = message


class GiftNameError(Exception):
    def __init__(self, message):
        self.message = message

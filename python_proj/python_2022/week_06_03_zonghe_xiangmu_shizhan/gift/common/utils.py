# coding:utf-8
import os
import time

from .error import NotPathError, NotFileError, FormatError


def timestamp_to_str(timestamp):
    time_obj = time.localtime(timestamp)
    time_str = time.strftime("%Y-%m-%d %H:%M:%S", time_obj)
    return time_str


def check_file(path):
    # 文件路径不存在
    if not os.path.exists(path):
        raise NotPathError('not found %s' % path)

    # 文件后缀名不是.json
    if not path.endswith('.json'):
        raise FormatError()

    # 不是文件
    if not os.path.isfile(path):
        raise NotFileError()

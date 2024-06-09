# coding:utf-8
import os

from .error import NotPathError, NotFileError, FormatError


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

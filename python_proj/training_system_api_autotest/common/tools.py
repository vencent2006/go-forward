import os


def get_project_path():
    """
    获取项目目录
    :return: 项目目录
    """
    # project_name的路径下不能存在重名的
    project_name = "training_system_api_autotest"
    file_path = os.path.dirname(__file__)
    return file_path[:file_path.find(project_name) + len(project_name)]


def sep(path, add_sep_before=False, add_sep_after=False):
    """
    系统分隔符
    :param path: 路径列表，类型为数组
    :param add_sep_before: 是否需要在拼接的路径前加一个分隔符
    :param add_sep_after: 是否需要在拼接的路径后加一个分隔符
    :return:
    """
    # path ["config","environment.yaml"]
    all_path = os.sep.join(path)
    if add_sep_before:
        all_path = os.sep + all_path
    if add_sep_after:
        all_path = all_path + os.sep

    return all_path


if __name__ == '__main__':
    print(get_project_path())
    print(sep(["config", "environment.yaml"], add_sep_before=True))

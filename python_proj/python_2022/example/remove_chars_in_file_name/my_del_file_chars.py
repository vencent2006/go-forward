import argparse
import os


def my_remove_chars(directory, strToDel):
    print("处理: %s\n" % directory)
    file_list = os.listdir(directory)
    for file_name in file_list:
        file_with_dir = os.path.join(directory, file_name)
        if os.path.isdir(file_with_dir):
            # 本身是个目录, 进到里面去
            # print("进入目录: %s\n" % file_with_dir)
            my_remove_chars(file_with_dir, strToDel)
            continue

        if file_name.find(strToDel) == -1:
            # 该文件名没有指定子串
            continue

        # 获取新的文件名
        new_file_name = file_name.replace(strToDel, '')

        # 构建文件的完整路径
        old_file_path = os.path.join(directory, file_name)
        new_file_path = os.path.join(directory, new_file_name)

        # 重命名文件
        os.rename(old_file_path, new_file_path)


parser = argparse.ArgumentParser()

parser.add_argument('-d', '--directory', help='目录名')
parser.add_argument('-s', '--string', help='要清除的字符串')

args = parser.parse_args()
directory = args.directory
strToDel = args.string
my_remove_chars(directory, strToDel)

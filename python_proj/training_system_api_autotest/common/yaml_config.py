import yaml
from tools import get_project_path, sep


class GetConf:
    def __init__(self):
        project_path = get_project_path()
        file_path = project_path + sep(["config", "environment.yaml"], add_sep_before=True)
        with open(file_path, "r") as env_file:
            self.env = yaml.load(env_file, yaml.FullLoader)
            # print(self.env)

    def get_username_password(self):
        return self.env["username"], self.env["password"]


if __name__ == '__main__':
    GetConf()
    print(GetConf().get_username_password())

# file_name = "/Users/mingzhe8/Documents/innovation/go_proj_exercise/vincent/go-forward/python_proj/training_system_api_autotest/config/environment.yaml"
# file = open(file_name, encoding="utf-8")
# try:
#     a = file.read()
#     print(a)
# except Exception as e:
#     print(e)
# finally:
#     file.close()

# with open(file_name, "r", encoding="utf-8") as file:
#     a = file.read()
#     print(a)
#
# with open(file_name, "r", encoding="utf-8") as file:
#     for i in file.readlines():
#         print("==========")
#         print(i.strip())

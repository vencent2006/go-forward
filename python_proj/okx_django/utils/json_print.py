import json

import constant


def print_pretty_json(json_str):
    if constant.print_json:
        print(json.dumps(json_str, indent=2))

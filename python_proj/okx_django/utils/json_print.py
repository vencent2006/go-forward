import json


def print_pretty_json(json_str):
    print(json.dumps(json_str, indent=2))

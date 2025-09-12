import json


def convert_to_json(input_string):
    """
    将 key:value 以分号分隔的字符串转换为 JSON 格式的字符串。

    :param input_string: 输入的 key:value 以分号分隔的字符串
    :return: 转换后的 JSON 格式的字符串
    """
    result = {}
    # 按分号分割字符串
    items = input_string.strip().split(';')
    for item in items:
        if ':' in item:
            # 按冒号分割键值对
            key, value = item.split(':', 1)
            result[key.strip()] = value.strip()
    # 将字典转换为 JSON 字符串
    return json.dumps(result, ensure_ascii=False, indent=2)


input_data = "uid:1216826604;mid:5186501358127905;oid:5186501358127905;summary:按技术面看，今天多空主力核心争夺点是B点（即3510点）。这位置将决定第二轮行情能不能真正启动。按原来技术结构看C5是3040点，那么以3040为第二轮行情的起动点的话，则3510点将是（III1）.......要是不过，将有（III2）调整，顺趋势的话下半年支持上移在的3186点。因此，我们要注意看技术面的变化！ "
json_output = convert_to_json(input_data)
print(json_output)

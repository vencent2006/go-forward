import json

jsonStr = "{\"code\":10000,\"msg\":\"succ\",\"data\":\"comment id 123\"}"
# 解析 JSON 字符串
parsed_data = json.loads(jsonStr)

# 提取 data 字段
data_field = parsed_data.get('data')

print(data_field)

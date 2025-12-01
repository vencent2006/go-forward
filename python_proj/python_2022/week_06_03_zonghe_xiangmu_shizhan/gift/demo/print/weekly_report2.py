# 把str中每行改成【稀缺装扮】   AI合照 - 开发
str_data = '''
AI合照	开发
'''


# 按行处理字符串
def format_weekly_report(text):
    # 分割成行并去除空行
    lines = [line.strip() for line in text.strip().split('\n')]
    formatted_lines = []

    for line in lines:
        # 使用制表符分割每行内容
        parts = line.split('\t')
        if len(parts) >= 2:
            # 确保每行都格式化为【项目名称】状态  进度  日期
            project_name = parts[0]
            status = parts[1]

            # 格式化行
            formatted_line = f"{project_name} - {status}"
            formatted_lines.append(formatted_line)

    return '\n'.join(formatted_lines)


# 处理并打印结果
if __name__ == "__main__":
    formatted_report = format_weekly_report(str_data)
    print("-- 格式化后的周报内容：\n\n")
    print(formatted_report)

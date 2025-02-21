import json

import requests

# 智谱 AI API 相关信息
APP_KEY = "cc43dde921a71125e14e7a506d6a97d8.JxMIqSibbF0neVLx"
API_URL = "https://open.bigmodel.cn/api/paas/v4/chat/completions"


# 模拟调用智谱 AI 的函数
def call_zhipu_ai(prompt):
    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {APP_KEY}"
    }
    data = {
        "model": "glm-4-flash",
        "messages": [
            {"role": "user", "content": prompt}
        ],
    }
    try:
        response = requests.post(API_URL, headers=headers, data=json.dumps(data))
        response.raise_for_status()
        result = response.json()
        message = result["choices"][0]["message"]["content"]
        return message
    except requests.RequestException as e:
        print(f"请求出错: {e}")
    except (KeyError, json.JSONDecodeError) as e:
        print(f"解析响应出错: {e}")
    return None


# Agent1: 读取并理解账本 1
def agent1(file_path):
    try:
        with open(file_path, 'r') as file:
            content = file.read()
            prompt = f"请读取并理解这个账本信息：{content}"
            result = call_zhipu_ai(prompt)
            print("Agent1 处理结果：")
            print(result)
            return content
    except FileNotFoundError:
        print(f"文件 {file_path} 未找到。")
        return None


# Agent2: 读取并理解账本 2
def agent2(file_path):
    try:
        with open(file_path, 'r') as file:
            content = file.read()
            prompt = f"请读取并理解这个账本信息：{content}"
            result = call_zhipu_ai(prompt)
            print("Agent2 处理结果：")
            print(result)
            return content
    except FileNotFoundError:
        print(f"文件 {file_path} 未找到。")
        return None


# Agent3: 对比前两个账本并整理输出最终结果
def agent3(content1, content2):
    if content1 and content2:
        prompt = f"请对比以下两个账本信息并整理输出最终结果：账本 1：{content1} 账本 2：{content2}"
        result = call_zhipu_ai(prompt)
        print("Agent3 处理结果：")
        print(result)
    else:
        print("无法进行对比，账本信息不完整。")


model = "glm-4-flash"


def generate_response(prompt):
    data = {
        "model": model,  # 可以根据实际情况修改模型名称
        "messages": [
            {"role": "user", "content": prompt}
        ]
    }
    try:
        response = requests.post(API_URL, headers=headers, data=json.dumps(data))
        response.raise_for_status()
        result = response.json()
        message = result["choices"][0]["message"]["content"]
        return message
    except requests.exceptions.RequestException as e:
        print(f"请求出错: {e}")
    except (KeyError, IndexError, json.JSONDecodeError):
        print("解析响应出错，请检查 API 返回格式。")
    return None


def main():
    print("开始聊天！输入 'quit' 退出。")
    while True:
        user_input = input("你: ")
        if user_input.lower() == 'quit':
            break
        response = generate_response(user_input)
        if response:
            print(f"{model}: {response}")


if __name__ == "__main__":
    import sys

    if len(sys.argv) != 3:
        print("用法: python script.py <账本1文件路径> <账本2文件路径>")
        sys.exit(1)

    file_path1 = sys.argv[1]
    file_path2 = sys.argv[2]

    content1 = agent1(file_path1)
    content2 = agent2(file_path2)
    agent3(content1, content2)
    main()

import json

import requests

# 请将以下值替换为你自己的 API 相关信息
APP_KEY = "cc43dde921a71125e14e7a506d6a97d8.JxMIqSibbF0neVLx"
API_URL = "https://open.bigmodel.cn/api/paas/v4/chat/completions"
headers = {
    "Content-Type": "application/json",
    "Authorization": f"Bearer {APP_KEY}"
}
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
    main()

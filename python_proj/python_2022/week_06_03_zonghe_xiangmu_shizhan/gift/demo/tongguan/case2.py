import json

import requests

# 智谱 AI API 相关信息
APP_KEY = "cc43dde921a71125e14e7a506d6a97d8.JxMIqSibbF0neVLx"
API_URL = "https://open.bigmodel.cn/api/paas/v4/chat/completions"

# 简单的本地知识库，模拟 RAG 的知识检索
knowledge_base = {
    "旅游推荐": "哇，旅游的话我觉得大理很不错呢！那里有美丽的洱海，风花雪月的美景一定会让你流连忘返。还有江南水乡乌镇，古色古香的建筑和宁静的氛围超棒哒。",
    "美食推荐": "好吃的可太多啦！重庆火锅一定要试试，热辣鲜香，超带感。还有广东早茶，各种点心精致又美味哟。",
    "健身建议": "健身的话，要记得循序渐进哦。可以先从一些简单的有氧运动开始，像慢跑、跳绳。平时也要注意饮食均衡，多吃蔬菜水果和蛋白质丰富的食物呀。"
}
model = "glm-4-flash"


def retrieve_knowledge(query):
    """从本地知识库中检索相关知识"""
    for key in knowledge_base:
        if key in query:
            return knowledge_base[key]
    return None


def chat_with_model(user_input):
    """调用智谱 AI 模型进行聊天"""
    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {APP_KEY}"
    }
    # 尝试从本地知识库检索信息
    retrieved_info = retrieve_knowledge(user_input)
    if retrieved_info:
        # 如果检索到信息，将其添加到用户输入中，模拟 RAG
        user_input = f"{user_input}\n相关补充信息：{retrieved_info}"
    print(user_input)

    data = {
        "model": model,  # 可以根据实际情况修改模型名称
        "messages": [
            {"role": "user", "content": user_input}
        ]
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


def main():
    print("嗨，我是刘亦菲数字人助手，有什么问题都可以问我哟！")
    while True:
        user_input = input("你: ")
        if user_input.lower() in ["退出", "quit", "exit"]:
            print("拜拜啦，祝你生活愉快！")
            break
        response = chat_with_model(user_input)
        if response:
            print("刘亦菲数字人助手:", response)


if __name__ == "__main__":
    main()

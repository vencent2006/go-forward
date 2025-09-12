from io import BytesIO

import opennsfw2 as n2
import requests


# 将自动下载预训练模型 open_nsfw_weights.h5 到 C:\Users\Administrator\.opennsfw2\weights
# pip install opennsfw2 requests

def predict_nsfw_from_url(image_url):
    """
    从给定的图片 URL 预测图片的 NSFW 概率。

    :param image_url: 图片的网络 URL
    :return: NSFW 概率值
    """
    try:
        # 发送请求获取图片内容
        response = requests.get(image_url, timeout=10)
        response.raise_for_status()  # 检查请求是否成功
        # 将响应内容转换为字节流
        image_stream = BytesIO(response.content)
        # 预测 NSFW 概率
        return n2.predict_image(image_stream)
    except requests.RequestException as e:
        print(f"请求图片 {image_url} 时出错: {e}")
        return None


# 指定网上图片地址列表
image_urls = [
    "https://i0.wp.com/wx4.sinaimg.cn/large/9df4d7b0ly8hpjjg20a97j20u00u0q5r.jpg",
    # "https://example.com/image2.png"
]

# 批量预测
for url in image_urls:
    nsfw_probability = predict_nsfw_from_url(url)
    if nsfw_probability is not None:
        print(f"URL: {url}, NSFW 概率: {nsfw_probability}")

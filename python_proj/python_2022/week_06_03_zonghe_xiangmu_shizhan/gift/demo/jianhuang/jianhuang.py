import os

import numpy as np
from tensorflow.keras.models import load_model
from tensorflow.keras.preprocessing import image

# 下载预训练模型文件 (第一次运行时会自动下载)
MODEL_URL = "https://github.com/GantMan/nsfw_model/releases/download/1.2.0/nsfw_mobilenet2.224x224.h5"
MODEL_PATH = "nsfw_model.h5"


def download_model():
    import urllib.request
    print("正在下载预训练模型...")
    urllib.request.urlretrieve(MODEL_URL, MODEL_PATH)
    print("模型下载完成")


def load_nsfw_model():
    """加载NSFW检测模型"""
    if not os.path.exists(MODEL_PATH):
        download_model()
    return load_model(MODEL_PATH)


def preprocess_image(img_path, target_size=(224, 224)):
    """预处理图片"""
    img = image.load_img(img_path, target_size=target_size)
    img_array = image.img_to_array(img)
    img_array = np.expand_dims(img_array, axis=0)
    img_array /= 255.0  # 归一化
    return img_array


def predict_nsfw_prob(image_path, model):
    """预测图片的NSFW概率"""
    try:
        # 预处理图片
        processed_img = preprocess_image(image_path)

        # 预测
        preds = model.predict(processed_img)

        # 返回各个类别的概率
        # 类别顺序: ['drawings', 'hentai', 'neutral', 'porn', 'sexy']
        return {
            'drawings': float(preds[0][0]),
            'hentai': float(preds[0][1]),
            'neutral': float(preds[0][2]),
            'porn': float(preds[0][3]),
            'sexy': float(preds[0][4]),
            'nsfw_score': float(preds[0][3] + preds[0][1] + preds[0][4])  # porn+hentai+sexy
        }
    except Exception as e:
        print(f"预测过程中出错: {e}")
        return None


def main():
    # 加载模型
    print("正在加载模型...")
    model = load_nsfw_model()
    print("模型加载完成")

    # 输入图片路径
    image_path = input("请输入图片路径: ").strip()

    if not os.path.exists(image_path):
        print("错误: 图片文件不存在")
        return

    # 进行预测
    results = predict_nsfw_prob(image_path, model)

    if results:
        print("\n检测结果:")
        print(f"绘画/卡通概率: {results['drawings']:.4f}")
        print(f"Hentai概率: {results['hentai']:.4f}")
        print(f"中性内容概率: {results['neutral']:.4f}")
        print(f"色情内容概率: {results['porn']:.4f}")
        print(f"性感内容概率: {results['sexy']:.4f}")
        print(f"\n综合NSFW评分: {results['nsfw_score']:.4f} (0-3范围)")

        # 给出判断建议
        if results['nsfw_score'] > 2.0:
            print("警告: 高概率包含不适宜内容!")
        elif results['nsfw_score'] > 1.0:
            print("注意: 可能包含敏感内容")
        else:
            print("安全: 未检测到明显不适宜内容")


if __name__ == "__main__":
    main()

import torch
from transformers import AutoTokenizer, AutoModelForSequenceClassification

# 加载预训练的多语言情感分析模型和分词器
MODEL_PATH = "/Users/vincent/.cache/huggingface/hub/models--nlptown--bert-base-multilingual-uncased-sentiment"
tokenizer = AutoTokenizer.from_pretrained(MODEL_PATH)
model = AutoModelForSequenceClassification.from_pretrained(MODEL_PATH)


def analyze_sentiment(text):
    """分析文本情感并返回1-10的整数评分"""
    # 预处理文本并进行模型推理
    inputs = tokenizer(text,
                       return_tensors="pt",
                       truncation=True,
                       max_length=512)

    with torch.no_grad():
        outputs = model(**inputs)

    # 计算概率分布
    probabilities = torch.softmax(outputs.logits, dim=1).numpy()[0]

    # 计算期望星级（1-5星范围）
    expected_star = sum((i + 1) * prob for i, prob in enumerate(probabilities))

    # 线性映射到1-10分范围并四舍五入
    score = round((expected_star - 1) * (9 / 4) + 1)

    # 确保分数在1-10范围内
    return max(1, min(10, score))


# 示例使用
if __name__ == "__main__":
    test_texts = [
        "这件衣服质量差到难以置信，第一次洗就褪色了！",  # 中文强烈负面
        "This is the best movie I've ever seen in my life!",  # 英文强烈正面
        "餐厅环境普通，菜品味道一般，没什么记忆点",  # 中文中性
        "The battery life is acceptable but the screen resolution could be better",  # 英文中性
        "完美无瑕的购物体验！快递准时，包装精致，客服专业",  # 中文强烈正面
        "Absolutely terrible customer service. Will never use this company again!",  # 英文强烈负面
        "操作界面设计合理，功能齐全但学习成本较高",  # 中文中性偏负面
        "The software works as advertised, though the price seems a bit high",  # 英文中性偏正面
        "彻底失望！承诺的功能根本没有实现",  # 中文强烈负面
        "这部小说让我重新思考人生的意义，强烈推荐给所有人",  # 中文强烈正面
        "Mediocre performance compared to other products in the same price range",  # 英文中性偏负面
        "机票价格合理，但航班延误处理非常不专业",  # 中文负面混合
        "Outstanding user experience with intuitive interface design",  # 英文强烈正面
        "咖啡口感中规中矩，甜点倒是意外地不错",  # 中文中性混合
        "The product arrived damaged but the replacement process was efficient",  # 英文中性混合
        "绝对物超所值！功能远超同价位竞品",  # 中文强烈正面
        "Worst purchase decision I've made this year",  # 英文强烈负面
        "剧情发展略显拖沓，但演员演技可圈可点",  # 中文中性混合
        "Reliable performance with occasional minor glitches",  # 英文中性偏正面
        "客服态度恶劣，彻底毁了这个品牌在我心中的形象",  # 中文强烈负面

    ]

    for text in test_texts:
        print(f"文本: {text}")
        print(f"情感评分: {analyze_sentiment(text)}\n")

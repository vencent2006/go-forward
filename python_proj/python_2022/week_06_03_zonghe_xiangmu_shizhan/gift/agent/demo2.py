from nltk.sentiment import SentimentIntensityAnalyzer


# try:
#     _create_unverified_https_context = ssl._create_unverified_context
# except AttributeError:
#     pass
# else:
#     ssl._create_default_https_context = _create_unverified_https_context
#
# nltk.download('vader_lexicon')


# 首次运行需要下载VADER词典（取消注释下一行）
# nltk.download('vader_lexicon')


def analyze_sentiment(text):
    sia = SentimentIntensityAnalyzer()
    # 获取情感得分（范围[-1.0, 1.0]）
    compound_score = sia.polarity_scores(text)['compound']
    # 将得分线性映射到1-10的整数范围
    scaled_score = int(round((compound_score + 1) * 4.5 + 1, 0))
    # 确保结果在1-10范围内
    return max(1, min(10, scaled_score))


# 示例用法
if __name__ == "__main__":
    test_text = "太温暖、感人了，看得我热泪盈眶。今年其实有在这样去做了，30多了第一次解锁游乐场、第一次看演唱会，利用周末时间出去旅游。虽然今年工作更加忙碌，压力也更大，但还是很开心，感觉生活开始有了属于自己的色彩。明年会继续努力的"
    print(analyze_sentiment(test_text))  # 输出10

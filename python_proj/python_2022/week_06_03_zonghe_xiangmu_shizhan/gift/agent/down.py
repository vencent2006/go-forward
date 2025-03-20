from transformers import AutoTokenizer, AutoModelForSequenceClassification

# 指定保存目录
MODEL_PATH = "./models/nlptown-bert-sentiment"

# 下载并保存模型&分词器（只需运行一次）
AutoTokenizer.from_pretrained("nlptown/bert-base-multilingual-uncased-sentiment").save_pretrained(MODEL_PATH)
AutoModelForSequenceClassification.from_pretrained("nlptown/bert-base-multilingual-uncased-sentiment").save_pretrained(
    MODEL_PATH)

import os

import opennsfw2 as n2


# 将自动下载预训练模型 open_nsfw_weights.h5 到 C:\Users\Administrator\.opennsfw2\weights
# pip install opennsfw2

def get_image_paths(directory, extensions):
    """
    获取指定目录下所有符合指定文件格式的文件路径。

    :param directory: 指定的目录路径
    :param extensions: 允许的文件扩展名列表，例如 ['.png', '.jpg', '.jpeg']
    :return: 符合条件的文件路径列表
    """
    local_image_paths = []
    for root, dirs, files in os.walk(directory):
        for file in files:
            if any(file.lower().endswith(ext) for ext in extensions):
                local_image_paths.append(os.path.join(root, file))
    return local_image_paths


# 指定目录和文件格式
image_directory = './pics'
image_extensions = ['.png', '.jpg', '.jpeg', '.gif']

# 获取符合条件的文件路径
image_paths = get_image_paths(image_directory, image_extensions)

# 批量预测
nsfw_probabilities = n2.predict_images(image_paths)
for index, value in enumerate(nsfw_probabilities):
    print(f"Index: {image_paths[index]}, Value: {value}")
# print(nsfw_probabilities)
# [0.16282965242862701, 0.8638442158699036]

import argparse  # 导入 argparse 库用于处理命令行参数
import concurrent.futures
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


# 解析命令行参数
parser = argparse.ArgumentParser(description='指定图片目录进行 NSFW 检测')
parser.add_argument('image_directory', type=str, help='包含图片的目录路径')
parser.add_argument('--max-workers', type=int, default=10, help='线程池的最大线程数，默认为 10')
args = parser.parse_args()

image_directory = args.image_directory
image_extensions = ['.png', '.jpg', '.jpeg', '.gif']

# 获取符合条件的文件路径
image_paths = get_image_paths(image_directory, image_extensions)


# 多线程批量预测
def predict_nsfw(image_path):
    try:
        probability = n2.predict_image(image_path)
        file_name = os.path.basename(image_path)
        print(f"{file_name}\t{probability:.2f}")  # 线程内直接输出结果
        return probability
    except Exception as e:
        print(f"处理文件 {os.path.basename(image_path)} 时出错: {e}")
        return None


# 使用命令行传入的线程数量
with concurrent.futures.ThreadPoolExecutor(max_workers=args.max_workers) as executor:
    # 不再需要收集结果列表，直接执行任务
    list(executor.map(predict_nsfw, image_paths))

# 命令执行例子
# python nsfw_pic_thread.py /path/to/your/image/directory --max-workers 20

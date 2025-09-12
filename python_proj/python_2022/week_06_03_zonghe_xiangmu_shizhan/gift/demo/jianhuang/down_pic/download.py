import argparse
import os
import urllib.parse

import requests


def download_images(txt_file_path, save_dir):
    # 检查保存目录是否存在，不存在则创建
    if not os.path.exists(save_dir):
        os.makedirs(save_dir)

    # 打开 txt 文件进行逐行读取
    with open(txt_file_path, 'r', encoding='utf-8') as file:
        for line in file:
            # 去除每行末尾的换行符
            image_url = line.strip()
            if image_url:
                try:
                    # 发送请求获取图片内容
                    response = requests.get(image_url, stream=True)
                    response.raise_for_status()

                    # 尝试从响应头中获取文件名
                    content_disposition = response.headers.get('Content-Disposition')
                    if content_disposition:
                        file_name = content_disposition.split('filename=')[-1].strip('"')
                    else:
                        # 从 URL 中提取文件名
                        url_path = urllib.parse.urlparse(image_url).path
                        file_name = os.path.basename(url_path)
                        if not file_name:
                            # 若无法获取文件名，抛出异常
                            raise ValueError(f"无法为 {image_url} 获取有效的文件名")

                    file_path = os.path.join(save_dir, file_name)

                    # 保存图片到本地文件
                    with open(file_path, 'wb') as img_file:
                        for chunk in response.iter_content(chunk_size=8192):
                            if chunk:
                                img_file.write(chunk)
                    print(f"成功下载: {image_url} 保存为: {file_name}")
                except Exception as e:
                    print(f"下载失败 {image_url}: {e}")
                    # 遇到错误，停止程序运行
                    raise


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='从 txt 文件中读取图片 URL 并下载到指定目录。')
    parser.add_argument('txt_file', help='包含图片 URL 的 txt 文件路径')
    parser.add_argument('save_dir', help='保存下载图片的目录路径')
    args = parser.parse_args()

    try:
        download_images(args.txt_file, args.save_dir)
    except Exception as e:
        print(f"程序因错误停止: {e}")

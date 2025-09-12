import os
from concurrent.futures import ThreadPoolExecutor, as_completed

import requests


def download_ts_file(number, base_url, save_dir):
    """
    下载单个TS文件

    Args:
        number: 文件编号
        base_url: 基础URL（不包含编号部分）
        save_dir: 保存目录
    """
    # 构建完整的下载URL
    url = base_url.replace("xxxx", f"{number}")

    # 构建保存文件名
    filename = f"v.f1598588_{number:04d}.ts"
    filepath = os.path.join(save_dir, filename)

    try:
        # 发送请求下载文件
        response = requests.get(url, stream=True, timeout=30)
        response.raise_for_status()  # 检查请求是否成功

        # 保存文件
        with open(filepath, 'wb') as f:
            for chunk in response.iter_content(chunk_size=8192):
                if chunk:
                    f.write(chunk)

        print(f"成功下载: {filename}")
        return True

    except Exception as e:
        print(f"下载失败 {filename}: {str(e)}")
        return False


def main():
    # 基础URL
    # base_url = "https://bstxvod.hokkj.cn/43bb9869vodtranscq1500037627/3502ef435145403696861081290/v.f1598588_xxxx.ts" # 加餐 day1
    # base_url = "https://bstxvod.hokkj.cn/43bb9869vodtranscq1500037627/53737ed93560136623208241277/v.f1598588_xxxx.ts"  # 加餐 day2
    # base_url = "https://bstxvod.hokkj.cn/43bb9869vodtranscq1500037627/12d580465145403695904737531/v.f1598588_xxxx.ts"  # day1
    base_url = "https://bstxvod.hokkj.cn/43bb9869vodtranscq1500037627/8d3a5f4b5145403691197461299/v.f1598588_xxxx.ts"  # day2

    # 创建保存目录
    save_dir = "ts_files"
    os.makedirs(save_dir, exist_ok=True)

    # 下载范围：1到1460
    start_num = 0
    # end_num = 1484 # 加餐 day1
    # end_num = 1513  # 加餐 day2
    # end_num = 1180  # day1
    end_num = 1600  # day2

    print(f"开始下载 {start_num} 到 {end_num} 的TS文件...")

    # 单线程下载（稳定但较慢）
    success_count = 0
    for i in range(start_num, end_num + 1):
        if download_ts_file(i, base_url, save_dir):
            success_count += 1

    print(f"下载完成！成功下载 {success_count}/{end_num - start_num + 1} 个文件")


def main_multithreaded():
    """
    多线程版本（下载更快，但可能对服务器压力较大）
    """
    # base_url = "https://bstxvod.hokkj.cn/43bb9869vodtranscq1500037627/3502ef435145403696861081290/v.f1598588_xxxx.ts" # day1
    base_url = "https://bstxvod.hokkj.cn/43bb9869vodtranscq1500037627/53737ed93560136623208241277/v.f1598588_270.ts"  # day2
    save_dir = "ts_files"
    os.makedirs(save_dir, exist_ok=True)

    start_num = 1
    end_num = 1460

    print(f"开始多线程下载 {start_num} 到 {end_num} 的TS文件...")

    # 使用线程池（最多8个线程）
    success_count = 0
    with ThreadPoolExecutor(max_workers=8) as executor:
        # 创建下载任务
        futures = {
            executor.submit(download_ts_file, i, base_url, save_dir): i
            for i in range(start_num, end_num + 1)
        }

        # 等待所有任务完成
        for future in as_completed(futures):
            if future.result():
                success_count += 1

    print(f"下载完成！成功下载 {success_count}/{end_num - start_num + 1} 个文件")


if __name__ == "__main__":
    # 选择下载方式：
    main()  # 单线程（稳定）
    # main_multithreaded()  # 多线程（快速）

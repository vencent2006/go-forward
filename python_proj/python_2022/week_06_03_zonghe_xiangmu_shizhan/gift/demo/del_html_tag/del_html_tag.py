import os
import re
import sys


def clean_text(text):
    """Remove HTML tags and &nbsp; from a string"""
    # Remove HTML tags
    text = re.sub(r'<.*?>', '', str(text))
    # Remove &nbsp; and other HTML entities
    text = re.sub(r'&nbsp;', ' ', text)
    text = re.sub(r'&[a-z]+;', ' ', text)  # 移除其他HTML实体如&amp;
    # 合并多个空格为一个
    # text = re.sub(r'\s+', ' ', text).strip()
    return text


def clean_text_file(input_file):
    # 生成输出文件名
    base_name, ext = os.path.splitext(input_file)
    output_file = f"{base_name}_2{ext}"

    try:
        # 读取输入文件内容
        with open(input_file, 'r', encoding='utf-8') as f:
            content = f.read()

        # 清理文本内容
        cleaned_content = clean_text(content)

        # 将清理后的内容写入输出文件
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write(cleaned_content)

        print(f"Cleaned data saved to {output_file}")
    except Exception as e:
        print(f"An error occurred: {e}")


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python del_html_tag.py <input_file>")
        sys.exit(1)

    input_file = sys.argv[1]
    if not os.path.exists(input_file):
        print(f"Error: File '{input_file}' not found")
        sys.exit(1)

    clean_text_file(input_file)

import json

from django.http import HttpResponse
from django.http import JsonResponse
from django.shortcuts import render

from blog.api_client import call_weibo_api


def get_weibo_data(request):
    """
    Django 视图方法，从请求中获取查询内容和用户 ID，
    调用微博 AIGC API 并返回处理后的数据。
    """
    if request.method == 'GET':
        query = request.GET.get('query')
        user_id = 'abc-user'
        if query and user_id:
            result = call_weibo_api(query, user_id)
            # print(f"API 响应: {json.dumps(result, ensure_ascii=False, indent=2)}")
            if result:
                return JsonResponse(result)
            else:
                return JsonResponse({"error": "请求微博 API 失败"}, status=500)
        return JsonResponse({"error": "缺少必要参数"}, status=400)
    return JsonResponse({"error": "只支持 GET 请求"}, status=405)


def convert_to_json(input_string):
    """
    将 key:value 以分号分隔的字符串转换为 JSON 格式的字符串。

    :param input_string: 输入的 key:value 以分号分隔的字符串
    :return: 转换后的 JSON 格式的字符串
    """
    result = {}
    # 按分号分割字符串
    items = input_string.strip().split(';')
    for item in items:
        if ':' in item:
            # 按冒号分割键值对
            key, value = item.split(':', 1)
            result[key.strip()] = value.strip()
    # 将字典转换为 JSON 字符串
    return json.dumps(result, ensure_ascii=False, indent=2)


def get_weibo_view(request):
    """
    Django 视图方法，从请求中获取查询内容和用户 ID，
    调用微博 AIGC API 并返回处理后的数据。
    """
    if request.method == 'GET':
        query = request.GET.get('query')
        user_id = 'abc-user'
        if query and user_id:
            result = call_weibo_api(query, user_id)
            print(f"API 响应: {json.dumps(result, ensure_ascii=False, indent=2)}")
            if result:
                data = result.get('data', {})
                outputs = data.get('outputs', {})
                res = outputs.get('res', '')
                # 将 Markdown 内容转换为 HTML
                # res_html = markdown.markdown(res)
                res_html = res.replace('\n', '<br>')
                kg_output = outputs.get('kg', {}).get('output', [])
                # 将 content 转换为 JSON 格式字符串
                content_list = [
                    convert_to_json(item.get('content', '').replace('\"', '').replace('"', ''))
                    for
                    item in kg_output]
                return render(request, 'blog/weibo_data.html',
                              {'res_html': res_html, 'content_list': content_list, 'question': query})
            else:
                return JsonResponse({"error": "请求微博 API 失败"}, status=500)
        return JsonResponse({"error": "缺少必要参数"}, status=400)
    return JsonResponse({"error": "只支持 GET 请求"}, status=405)


# Create your views here.
def hello_world(request):
    return HttpResponse("Hello, world")


def picture(request):
    return render(request, 'blog/picture.html')


def get_picture(request):
    index = int(request.GET.get('index', 0))
    # 这里需要替换为你的图片列表
    picture_list = [
        'https://example.com/picture1.jpg',
        'https://example.com/picture2.jpg',
        'https://example.com/picture3.jpg',
    ]
    if 0 <= index < len(picture_list):
        url = picture_list[index]
    else:
        url = ''
    return JsonResponse({'url': url})

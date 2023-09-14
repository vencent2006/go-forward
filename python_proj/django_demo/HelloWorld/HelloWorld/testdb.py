from django.http import HttpResponse
from TestModel.models import entLocalWhitelist


# 数据库操作
def testdb(request):
    # 初始化
    response = ""
    response1 = ""
    li = entLocalWhitelist.objects.all()
    # 输出所有数据
    for var in li:
        response1 += str(var.uid) + " "
    response = response1
    return HttpResponse("<p>" + response + "</p>")

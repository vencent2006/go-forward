from django.http import HttpResponse
from django.shortcuts import render


def hello(request):
    return HttpResponse('Hello World!')


def runoob(request):
    context = {}
    context['hello'] = 'Hello vincent!'
    return render(request, 'runoob.html', context)


def runoob2(request):
    views_name = 'Hello lucas!'
    return render(request, 'runoob2.html', {"name": views_name})


def runoob3(request):
    views_list = ["菜鸟教程1", "菜鸟教程2", "菜鸟教程3"]
    return render(request, "runoob3.html", {"views_list": views_list})

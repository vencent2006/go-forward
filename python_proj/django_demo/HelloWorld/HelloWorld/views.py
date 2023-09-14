from django.shortcuts import render

def runoob(request):
    name =0
    return render(request, "runoob.html", {"name": name})

from django.urls import path

import blog.views

urlpatterns = [
    path('hello_world', blog.views.hello_world),
    path('wb_zhishi', blog.views.get_weibo_data),
    path('wb_zhishi_view', blog.views.get_weibo_view),
    path('get-picture/', blog.views.get_picture, name='get-picture'),
]

from django.urls import path
from django.conf.urls import url


from . import views, testdb, search

urlpatterns = [
    path('runoob/', views.runoob),
    path('testdb/', testdb.testdb),
    url(r'^search-form/$', search.search_form),
    url(r'^search/$', search.search),
]

from django.urls import path

from . import views
from . import testdb

urlpatterns = [
    path('runoob/', views.runoob),
    path('testdb/', testdb.testdb),
]

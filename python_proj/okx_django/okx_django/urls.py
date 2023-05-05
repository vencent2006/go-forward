"""okx_django URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.2/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path

import controller

urlpatterns = [
    path('admin/', admin.site.urls),
    # market
    path('market/tickers/', controller.get_tickers),
    path('market/ticker/', controller.get_ticker),
    # account balance
    path('account/balance/all', controller.get_all_account_balance),
    path('account/balance/one', controller.get_one_account_balance),
    # public data
    path('public/instruments', controller.get_instruments),
]

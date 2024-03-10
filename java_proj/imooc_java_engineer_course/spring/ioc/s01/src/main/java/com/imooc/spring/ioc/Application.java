package com.imooc.spring.ioc;

import com.imooc.spring.ioc.entity.Apple;
import com.imooc.spring.ioc.entity.Child;

public class Application {
    public static void main(String[] args) {
        Apple apple1 = new Apple("红富士", "红色", "欧洲");
        Apple apple2 = new Apple("青苹果", "绿色", "中亚");
        Apple apple3 = new Apple("金帅", "黄色", "中国");

        Child lily = new Child("莉莉", apple1);
        Child andy = new Child("安迪", apple2);
        Child luna = new Child("露娜", apple3);

        lily.eat();
        andy.eat();
        luna.eat();
    }
}

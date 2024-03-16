package com.imooc.spring.ioc;

import com.imooc.spring.ioc.entity.Apple;
import com.imooc.spring.ioc.entity.Child;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class SpringApplication {
    public static void main(String[] args) {
        // 初始化ioc容器
        ApplicationContext context = new ClassPathXmlApplicationContext("classpath:applicationContext.xml");
        Apple sweetApple = context.getBean("sweetApple", Apple.class);
        System.out.println(sweetApple.getTitle());
        Child lily = context.getBean("lily", Child.class);
        lily.eat();
        Child andy = context.getBean("andy", Child.class);
        andy.eat();
        Child luna = context.getBean("luna", Child.class);
        luna.eat();
    }
}

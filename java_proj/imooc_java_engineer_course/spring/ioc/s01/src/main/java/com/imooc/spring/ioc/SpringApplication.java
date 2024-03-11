package com.imooc.spring.ioc;

import com.imooc.spring.ioc.entity.Apple;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class SpringApplication {
    public static void main(String[] args) {
        // 初始化ioc容器
        ApplicationContext context = new ClassPathXmlApplicationContext("classpath:applicationContext.xml");
        Apple sweetApple = context.getBean("sweetApple", Apple.class);
        System.out.println(sweetApple.getTitle());
    }
}

package com.jingdianjichi.user.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.concurrent.LinkedBlockingDeque;
import java.util.concurrent.ThreadFactory;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

@Configuration
public class ThreadPoolConfig {
    @Bean(name = "mailThreadPool")
    public ThreadPoolExecutor getMailThreadPool() {
        ThreadFactory threadFactory = new CustomNameThreadFactory("mail");
        return new ThreadPoolExecutor(
                20,
                50,
                5,
                TimeUnit.SECONDS,
                new LinkedBlockingDeque<>(50),
                threadFactory,// 工厂函数
                new ThreadPoolExecutor.DiscardPolicy()
        );
    }
}

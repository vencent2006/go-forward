package com.vincent.utils;

import lombok.extern.slf4j.Slf4j;
import org.springframework.context.annotation.Configuration;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.annotation.Scheduled;

import java.time.LocalDateTime;

// @Configuration // 1. 标记配置类，使得springboot容器可以扫描到
// @EnableScheduling // 2. 开启定时任务
@Slf4j
public class MyTask {
    // 3. 添加一个任务，并且注明任务的运行表达式
    @Scheduled(cron = "*/5 * * * * ?") // 每隔5秒
    public void publishMsg(){
        log.warn("开始执行任务:" + LocalDateTime.now());
    }
}

package com.vincent.utils;

import lombok.extern.slf4j.Slf4j;
import org.springframework.scheduling.annotation.Async;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.stereotype.Component;

import java.time.LocalDateTime;

@Component
@EnableAsync
@Slf4j
public class MyAsyncTask {
    @Async
    public void publicMsg() {
        try {
            Thread.sleep(5000);
            log.warn("异步任务处理完毕");
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}

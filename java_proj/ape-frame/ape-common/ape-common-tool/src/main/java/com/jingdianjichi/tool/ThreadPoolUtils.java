package com.jingdianjichi.tool;

import lombok.extern.slf4j.Slf4j;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.TimeUnit;

@Slf4j
public class ThreadPoolUtils {
    // 私有构造方法
    private ThreadPoolUtils() {
    }

    public static void shutDownPool(ExecutorService pool, int shutdownTimeout, int shutdownNowTimeout, TimeUnit timeUnit) {
        pool.shutdown();
        try {
            if (!pool.awaitTermination(shutdownTimeout, timeUnit)) {
                pool.shutdownNow(); // 取消正在执行的任务
                if (!pool.awaitTermination(shutdownNowTimeout, timeUnit)) {
                    log.error("ThreadPoolUtils.shutDownPool.error");
                }
            }
        } catch (InterruptedException e) {
            log.error("ThreadPoolUtils.shutDownPool.interrupt.error:{}", e.getMessage(), e);
            pool.shutdownNow();// 取消正在执行的任务
            Thread.currentThread().interrupt();// 保留中断状态
        }
    }
}

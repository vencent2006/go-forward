package com.jingdianjichi.user;

import lombok.extern.slf4j.Slf4j;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

@SpringBootTest(classes = UserApplication.class, webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@RunWith(SpringRunner.class)
@Slf4j
public class ThreadPoolShutDownTest {
    @Test
    public void testShutdown() throws Exception {
        ExecutorService executorService = Executors.newFixedThreadPool(10);
        for (int i = 0; i < 1000; i++) {
            executorService.execute(new TaskShutDownPool());
        }
        Thread.sleep(1000); // 等待1秒
        // log.info("testShutdown.status:{}, before", executorService.isShutdown());
        log.info("testShutdown.status:{}, before", executorService.isTerminated());
        // executorService.shutdown(); // 关闭线程池
        executorService.shutdownNow(); // 立即关闭线程池
        // log.info("testShutdown.status:{}, after", executorService.isShutdown());
        // log.info("testShutdown.status:{}, after", executorService.isTerminated());
        log.info("testShutdown.status:{}, after", executorService.awaitTermination(10L, TimeUnit.SECONDS));
        Thread.sleep(500); // 等待0.5秒
        log.info("testShutdown.shutdown");
        // executorService.execute(new TaskShutDownPool());
    }

    class TaskShutDownPool implements Runnable {
        @Override
        public void run() {
            try {
                Thread.sleep(500);
                log.info("TaskShutDownPool.{}", Thread.currentThread().getName());
            } catch (InterruptedException e) {
                log.info("TaskShutDownPool.interrupted:{}", e.getMessage(), e);
            }
        }
    }
}

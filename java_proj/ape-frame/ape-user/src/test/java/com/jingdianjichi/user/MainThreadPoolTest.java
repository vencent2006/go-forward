package com.jingdianjichi.user;

import com.jingdianjichi.tool.CompletableFutureUtils;
import lombok.extern.slf4j.Slf4j;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import javax.annotation.Resource;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.FutureTask;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

// classes 表示启动类
// webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT 表示随机端口启动应用
@SpringBootTest(classes = UserApplication.class, webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@RunWith(SpringRunner.class) // 表示启动这个单元测试类（单元测试类是不能够运行的），需要传递一个参数，必须是SpringRunner的实例类型
@Slf4j
public class MainThreadPoolTest {
    @Resource(name = "mailThreadPool")
    private ThreadPoolExecutor mailThreadPool;

    @Test
    public void test() {
        for (int i = 0; i < 10; i++) {
            mailThreadPool.submit(new Runnable() {
                @Override
                public void run() {
                    log.info("当前时间 {}", System.currentTimeMillis());
                }
            });
        }
    }

    @Test
    public void testFuture() {
        List<FutureTask<String>> futureTaskList = new ArrayList<>();

        FutureTask futureTask1 = new FutureTask<>(() -> {
            Thread.sleep(2000);// 模拟超时操作
            return "鸡翅";
        });
        FutureTask futureTask2 = new FutureTask<>(() -> {
            return "经典";
        });

        futureTaskList.add(futureTask1);
        futureTaskList.add(futureTask2);

        mailThreadPool.submit(futureTask1);
        mailThreadPool.submit(futureTask2);

        for (int i = 0; i < futureTaskList.size(); i++) {
            String name = CompletableFutureUtils.getResult(futureTaskList.get(i),
                    1, TimeUnit.SECONDS, "默认值", log);
            log.info("testFuture: {}", name);
        }
    }
}

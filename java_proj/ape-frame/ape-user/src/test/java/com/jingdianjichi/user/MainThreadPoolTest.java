package com.jingdianjichi.user;

import lombok.extern.slf4j.Slf4j;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import javax.annotation.Resource;
import java.util.concurrent.ThreadPoolExecutor;

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
}

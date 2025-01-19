package com.jingdianjichi.user.controller;

import com.jingdianjichi.redis.util.RedisShareLockUtil;
import com.jingdianjichi.redis.util.RedisUtil;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@Slf4j
public class TestController {
    @Autowired
    private RedisUtil redisUtil;

    @Autowired
    private RedisShareLockUtil redisShareLockUtil;

    @GetMapping("/test")
    public String test() {
        return "test";
    }

    @GetMapping("/testRedis")
    public void testRedis() {
        redisUtil.set("name", "鸡翅");
    }

    @GetMapping("/testRedisLock")
    public void testRedisLock() {
        boolean jichi = redisShareLockUtil.lock("jichi", "1231231", 100000L);
        System.out.println(jichi);
    }

    @GetMapping("/testLog")
    public void testLog() {
        long startTime = System.currentTimeMillis();
        for (int i = 0; i < 100000; i++) {
            log.info("这是{}个日志", i);
        }
        long endTime = System.currentTimeMillis();
        log.info("耗时：" + (endTime - startTime) + "ms");
    }
}

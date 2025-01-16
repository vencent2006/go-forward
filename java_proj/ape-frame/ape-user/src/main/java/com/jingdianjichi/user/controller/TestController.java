package com.jingdianjichi.user.controller;

import com.jingdianjichi.redis.util.RedisUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TestController {
    @Autowired
    private RedisUtil redisUtil;

    @GetMapping("/test")
    public String test() {
        return "test";
    }

    @GetMapping("/testRedis")
    public void testRedis() {
        redisUtil.set("name", "鸡翅");
    }
}

package com.jingdianjichi.user.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TestController {
    @Autowired
    private RedisTemplate redisTemplate;

    @GetMapping("/test")
    public String test() {
        return "test";
    }

    @GetMapping("/testRedis")
    public void testRedis() {
        redisTemplate.opsForValue().set("name", "鸡翅");
    }
}

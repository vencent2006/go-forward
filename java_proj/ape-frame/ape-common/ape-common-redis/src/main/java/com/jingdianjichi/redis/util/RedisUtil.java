package com.jingdianjichi.redis.util;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Component;

import java.util.concurrent.TimeUnit;

@Component
public class RedisUtil {
    @Autowired
    private RedisTemplate redisTemplate;

    public void set(String key, String value) {
        redisTemplate.opsForValue().set(key, value);
    }

    public boolean setNx(String key, String value, long timeout, TimeUnit timeunit) {
        return redisTemplate.opsForValue().setIfAbsent(key, value, timeout, timeunit);
    }

    public String get(String key) {
        return (String) redisTemplate.opsForValue().get(key);
    }

    public boolean del(String key) {
        return redisTemplate.delete(key);
    }
}

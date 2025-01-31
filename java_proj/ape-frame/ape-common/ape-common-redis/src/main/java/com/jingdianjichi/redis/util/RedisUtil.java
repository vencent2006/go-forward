package com.jingdianjichi.redis.util;

import com.alibaba.fastjson.JSON;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.ClassPathResource;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.script.DefaultRedisScript;
import org.springframework.scripting.support.ResourceScriptSource;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.util.ArrayList;
import java.util.List;
import java.util.Set;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;
import java.util.stream.Stream;

@Component
public class RedisUtil {
    @Autowired
    private RedisTemplate redisTemplate;


    private static final String CACHE_KEY_SEPARATOR = ".";

    private DefaultRedisScript<Boolean> casScript;

    @PostConstruct // 表示在构造方法执行之后执行
    public void init() {
        // 加载lua脚本
        casScript = new DefaultRedisScript<>();
        casScript.setResultType(Boolean.class);
        // 脚本位置
        casScript.setScriptSource(new ResourceScriptSource(new ClassPathResource("compareAndSet.lua")));
        System.out.println(JSON.toJSON(casScript));
    }

    public Boolean compareAndSet(String key, Long oldValue, Long newValue) {
        List<String> keys = new ArrayList();
        keys.add(key);
        return (Boolean) redisTemplate.execute(casScript, keys, oldValue, newValue);
    }

    /**
     * 构建缓存key
     */
    public String buildKey(String... strObjs) {
        return Stream.of(strObjs).collect(Collectors.joining(CACHE_KEY_SEPARATOR));
    }

    /**
     * 是否存在key
     */
    public boolean exist(String key) {
        return redisTemplate.hasKey(key);
    }

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


    public Boolean zAdd(String key, String value, Long score) {
        return redisTemplate.opsForZSet().add(key, value, Double.valueOf(String.valueOf(score)));
    }

    public Long countZset(String key) {
        return redisTemplate.opsForZSet().size(key);
    }

    public Set<String> rangeZset(String key, long start, long end) {
        return redisTemplate.opsForZSet().range(key, start, end);
    }

    public Long removeZset(String key, Object value) {
        return redisTemplate.opsForZSet().remove(key, value);
    }

    public void removeZsetList(String key, Set<String> value) {
        value.stream().forEach((val) -> redisTemplate.opsForZSet().remove(key, val));
    }

    public Double score(String key, Object value) {
        return redisTemplate.opsForZSet().score(key, value);
    }

    public Set<String> rangeByScore(String key, long start, long end) {
        return redisTemplate.opsForZSet().rangeByScore(key, Double.valueOf(String.valueOf(start)), Double.valueOf(String.valueOf(end)));
    }

    public Object addScore(String key, Object obj, double score) {
        return redisTemplate.opsForZSet().incrementScore(key, obj, score);
    }

    public Object rank(String key, Object obj) {
        return redisTemplate.opsForZSet().rank(key, obj);
    }


}

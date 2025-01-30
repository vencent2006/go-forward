package com.jingdianjichi.user.controller;

import com.jingdianjichi.redis.util.CacheUtil;
import com.jingdianjichi.redis.util.RedisShareLockUtil;
import com.jingdianjichi.redis.util.RedisUtil;
import com.jingdianjichi.user.entity.po.SysUser;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Map;

@RestController
@Slf4j
public class TestController {
    @Autowired
    private RedisUtil redisUtil;

    @Autowired
    private RedisShareLockUtil redisShareLockUtil;

    @Resource
    private CacheUtil cacheUtil;

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

    @GetMapping("/testLocalCache")
    public void testLocalCache() {
        // list<long> skuId
        // 价格 商品名称 限购 批量
        // Map<long, skuInfo>
        // 反射
        List<Long> skuIdList = new ArrayList<>();
        cacheUtil.getResult(skuIdList, "skuInfo.skuName", SkuInfo.class, (list) -> {
            Map<Long, SkuInfo> skuInfoMap = getSkuName(skuIdList);
            return skuInfoMap;
        });
        cacheUtil.getResult(skuIdList, "skuInfo.skuPrice", SkuPriceInfo.class, (list) -> {
            Map<Long, SkuPriceInfo> skuInfoMap = getSkuPrice(skuIdList);
            return skuInfoMap;
        });
    }

    public Map<Long, SkuInfo> getSkuName(List<Long> skuIdList) {
        return Collections.emptyMap();
    }

    public Map<Long, SkuPriceInfo> getSkuPrice(List<Long> skuIdList) {
        return Collections.emptyMap();
    }

    class SkuInfo {
        private Long id;
        private String name;
    }

    class SkuPriceInfo {
        private Long id;
        private Long price;
    }

    @PostMapping("/testQuery")
    public void testQuery(@RequestBody SysUser sysUser) throws Exception {
        // 2022-12-18 22:13:14
        System.out.println(sysUser);
    }
}

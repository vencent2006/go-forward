package com.jingdianjichi.redis.util;

import com.alibaba.fastjson.JSON;
import com.google.common.cache.Cache;
import com.google.common.cache.CacheBuilder;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang.StringUtils;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.util.CollectionUtils;

import java.util.*;
import java.util.concurrent.TimeUnit;
import java.util.function.Function;

@Component
@Slf4j
public class CacheUtil<K, V> {
    // 是否开启缓存
    @Value("${guava.cache.switch}")
    public Boolean switchCache;

    // 初始化本地缓存
    private Cache<String, String> localCache = CacheBuilder.newBuilder()
            .maximumSize(5000)
            .expireAfterWrite(3, TimeUnit.SECONDS)
            .build();

    public Map<K, V> getResult(List<K> skuIdList, String cachePrefix, Class<V> clazz, Function<List<K>, Map<K, V>> function) {
        // 参数校验
        if (CollectionUtils.isEmpty(skuIdList)) {
            return Collections.emptyMap();
        }
        // 声明一个结果类
        Map<K, V> resultMap = new HashMap<>(16);
        // 如果不开启缓存，直接调用方法
        if (!switchCache) {
            resultMap = function.apply(skuIdList);
            return resultMap;
        }
        // 声明本地缓存没有的id集合
        List<K> noCacheIdList = new ArrayList<>();
        // 遍历id集合
        for (K id : skuIdList) {
            String cacheKey = cachePrefix + "_" + id;
            String content = localCache.getIfPresent(cacheKey);
            if (StringUtils.isNotBlank(content)) {
                V v = JSON.parseObject(content, clazz);
                resultMap.put(id, v);
            } else {
                noCacheIdList.add(id);
            }
        }
        // 本地缓存没有的id集合为空，直接返回
        if (CollectionUtils.isEmpty(noCacheIdList)) {
            return resultMap;
        }
        Map<K, V> noCacheResultMap = function.apply(noCacheIdList);
        if (noCacheResultMap == null || noCacheResultMap.isEmpty()) {
            return resultMap;
        }
        for (Map.Entry<K, V> entry : noCacheResultMap.entrySet()) {
            K id = entry.getKey();
            V v = entry.getValue();
            resultMap.put(id, v);
            String cacheKey = cachePrefix + "_" + id;
            localCache.put(cacheKey, JSON.toJSONString(v));
        }
        return resultMap;
    }
}

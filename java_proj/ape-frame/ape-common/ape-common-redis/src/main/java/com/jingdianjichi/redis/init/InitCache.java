package com.jingdianjichi.redis.init;

import com.jingdianjichi.redis.util.SpringContextUtil;
import org.springframework.boot.CommandLineRunner;
import org.springframework.context.ApplicationContext;
import org.springframework.stereotype.Component;

import java.util.Map;

@Component
public class InitCache implements CommandLineRunner {
    @Override
    public void run(String... args) throws Exception {
        // 我要知道哪些缓存需要预热
        ApplicationContext applicationContext = SpringContextUtil.getApplicationContext();
        Map<String, AbstractCache> beanMap = applicationContext.getBeansOfType(AbstractCache.class);

        // 调用init方法
        if (beanMap.isEmpty()) {
            // 没有实现类
            return;
        }
        for (Map.Entry<String, AbstractCache> entry : beanMap.entrySet()) {
            // AbstractCache abstractCache = (AbstractCache) SpringContextUtil.getBean(entry.getValue().getClass());
            // abstractCache.initCache();
            AbstractCache value = entry.getValue();
            value.initCache();
        }
    }
}

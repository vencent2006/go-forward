package com.jingdianjichi.user.init;

import com.alibaba.fastjson.JSON;
import com.jingdianjichi.user.designPattern.builderPattern.demo.SkuDO;
import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.ApplicationListener;
import org.springframework.stereotype.Component;

import java.util.HashMap;
import java.util.Map;

@Component
@Slf4j
public class ApplicationInit implements ApplicationListener<ApplicationReadyEvent> {
    Map<String, initFunction> initFunctionMap = new HashMap<>();

    {
        initFunctionMap.put("预热fastjson", this::initFastJson);
    }

    @Override
    public void onApplicationEvent(ApplicationReadyEvent event) {
        // 预热httpclient
        // 预热json
        // 预热一个线程池
        initFunctionMap.forEach((desc, function) -> {
            try {
                long start = System.currentTimeMillis();
                function.invoke();
                long end = System.currentTimeMillis();
                log.info("ApplicationInit:{}, costTime:{}ms", desc, end - start);
            } catch (Exception e) {
                log.error("ApplicationInit{}.error", desc);
            }
        });
    }

    private void initFastJson() {
        SkuDO skuDO = new SkuDO();
        skuDO.setSkuId(1L);
        skuDO.setSkuName("苹果");
        String s = JSON.toJSONString(skuDO);
        System.out.println(s);
        JSON.parseObject(s, SkuDO.class);
    }

    interface initFunction {
        void invoke();
    }
}

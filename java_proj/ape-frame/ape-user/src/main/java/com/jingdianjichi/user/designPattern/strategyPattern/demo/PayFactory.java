package com.jingdianjichi.user.designPattern.strategyPattern.demo;

import org.springframework.beans.factory.InitializingBean;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Component
public class PayFactory implements InitializingBean {

    @Resource
    private List<PayHandler> payHandlerList;// 将PayHandler的所有实现类都注入到这个集合中

    private Map<PayChannelEnum, PayHandler> handlerMap = new HashMap<>();

    public PayHandler getHandlerByCode(int code) {
        PayChannelEnum payChannelEnum = PayChannelEnum.getByCode(code);
        return handlerMap.get(payChannelEnum);
    }

    @Override
    public void afterPropertiesSet() throws Exception {
        for (PayHandler payHandler : payHandlerList) {
            handlerMap.put(payHandler.getChannel(), payHandler);
        }
    }
}

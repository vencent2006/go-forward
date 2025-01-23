package com.jingdianjichi.user.designPattern.strategyPattern.demo;

import org.springframework.stereotype.Component;

import javax.annotation.Resource;

@Component
public class PayHandlerDemo {

    @Resource
    private PayFactory payFactory;

    public void dealPay(int code) {
        PayHandler payHandler = payFactory.getHandlerByCode(code);
        payHandler.dealPay();
    }

}

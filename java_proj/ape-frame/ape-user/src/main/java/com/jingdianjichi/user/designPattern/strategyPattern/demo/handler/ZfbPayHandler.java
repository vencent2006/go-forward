package com.jingdianjichi.user.designPattern.strategyPattern.demo.handler;

import com.jingdianjichi.user.designPattern.strategyPattern.demo.PayChannelEnum;
import com.jingdianjichi.user.designPattern.strategyPattern.demo.PayHandler;
import org.springframework.stereotype.Component;

@Component
public class ZfbPayHandler implements PayHandler {

    @Override
    public PayChannelEnum getChannel() {
        return PayChannelEnum.ZFB_PAY;
    }

    @Override
    public void dealPay() {
        System.out.println("支付宝支付");
    }
}

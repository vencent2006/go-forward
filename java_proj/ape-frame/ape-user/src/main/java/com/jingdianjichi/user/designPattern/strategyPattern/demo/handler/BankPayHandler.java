package com.jingdianjichi.user.designPattern.strategyPattern.demo.handler;

import com.jingdianjichi.user.designPattern.strategyPattern.demo.PayChannelEnum;
import com.jingdianjichi.user.designPattern.strategyPattern.demo.PayHandler;
import org.springframework.stereotype.Component;

@Component
public class BankPayHandler implements PayHandler {

    @Override
    public PayChannelEnum getChannel() {
        return PayChannelEnum.BANK_PAY;
    }

    @Override
    public void dealPay() {
        System.out.println("银行卡支付");
    }

}

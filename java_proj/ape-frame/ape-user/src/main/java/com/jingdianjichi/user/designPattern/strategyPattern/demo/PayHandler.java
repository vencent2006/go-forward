package com.jingdianjichi.user.designPattern.strategyPattern.demo;

public interface PayHandler {

    PayChannelEnum getChannel();

    void dealPay();

}

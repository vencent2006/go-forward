package com.imooc.controller;

import org.springframework.web.bind.annotation.RestController;

@RestController
public class BaseController {
    // 切换大小写：command + shift + u
    public static final Integer COMMENT_PAGE_SIZE = 10;
    public static final Integer PAGE_SIZE = 20;

    public static final String FOODIE_SHOPCAT = "shopcart"; // 前端购物车的cookie的key

    // 支付中心的调用地址
    String paymentUrl = "http://payment.t.mukewang.com/foodie-payment/payment/createMerchantOrder";

    // 微信支付成功 -> 支付中心 -> 天天吃货平台
    //                      |-> 回调通知的url
    String payReturnUrl = "http://localhost:8088/orders/notifyMerchantOrderPaid";
}

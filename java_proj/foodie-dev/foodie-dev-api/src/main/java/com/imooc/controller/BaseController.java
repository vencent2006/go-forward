package com.imooc.controller;

import com.imooc.pojo.Orders;
import com.imooc.service.center.MyOrdersService;
import com.imooc.utils.IMOOCJSONResult;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RestController;

import java.io.File;

@RestController
public class BaseController {
    // 切换大小写：command + shift + u
    public static final Integer COMMON_PAGE_SIZE = 10;
    public static final Integer PAGE_SIZE = 20;

    public static final String FOODIE_SHOPCAT = "shopcart"; // 前端购物车的cookie的key

    @Autowired
    private MyOrdersService myOrdersService;

    // 支付中心创建订单的调用地址
    String paymentUrl = "http://payment.t.mukewang.com/foodie-payment/payment/createMerchantOrder";

    // 支付中心查询订单的调用地址
    // post的params格式: merchantOrderId={merchantOrderId}&merchantUserId={merchantUserId}
    String paymentQueryMerchantOrderUrl = "http://payment.t.mukewang.com/foodie-payment/payment/getPaymentCenterOrderInfo";

    // 微信支付成功 -> 支付中心 -> 天天吃货平台
    //                      |-> 回调通知的url
    String payReturnUrl = "http://localhost:8088/orders/notifyMerchantOrderPaid";

    // 用户上传头像的位置 利用File.separator兼容Windows和类linux
    public static final String IMAGE_USER_FACE_LOCATION = File.separator + "workspaces" +
                                                            File.separator + "images" +
                                                            File.separator + "foodie" +
                                                            File.separator + "faces";


    /**
     * 检查用户订单
     * @param userId 用户id
     * @param orderId 订单id
     * @return IMOOCJSONResult
     */
    protected IMOOCJSONResult checkUserOrder(String userId, String orderId) {

        Orders order = myOrdersService.queryMyOrder(userId, orderId);
        if (null == order) {
            return IMOOCJSONResult.errorMsg("订单不存在");
        }

        return IMOOCJSONResult.ok(order);
    }
}

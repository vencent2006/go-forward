package com.imooc.service;

import com.imooc.pojo.bo.SubmitOrderBO;

public interface OrderService {
    /**
     * 创建订单
     * @param submitOrderBO 客户端请求的BO
     * @return String orderId（主订单的订单id）
     */
    public String createOrder(SubmitOrderBO submitOrderBO);


    /**
     * 修改订单状态
     * @param orderId 订单id
     * @param orderStatus 订单状态
     */
    public void updateOrderStatus(String orderId, Integer orderStatus);
}

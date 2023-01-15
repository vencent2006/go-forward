package com.imooc.service;

import com.imooc.utils.PagedGridResult;

public interface MyOrdersService {

    /**
     * 查询我的订单列表
     * @param userId 用户id
     * @param orderStatus 订单状态
     * @param page 第几页
     * @param pageSize 每页数量
     * @return 分页结果
     */
    public PagedGridResult queryMyOrders(String userId, Integer orderStatus, Integer page, Integer pageSize);
}

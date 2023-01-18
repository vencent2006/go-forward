package com.imooc.service.center;

import com.imooc.pojo.Orders;
import com.imooc.pojo.vo.OrderStatusCountsVO;
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

    /**
     * 订单状态 --> 商家发货
     * @param orderId 订单id
     */
    public void updateDeliverOrderStatus(String orderId);

    /**
     * 查询我的订单
     * @param userId 用户id
     * @param orderId 订单id
     * @return Orders 订单信息
     */
    public Orders queryMyOrder(String userId, String orderId);


    /**
     * 订单状态 --> 用户确认收货
     * @param orderId 订单id
     * @return boolean true: 成功; false: 失败
     */
    public boolean updateReceiveOrderStatus(String orderId);


    /**
     * 删除订单（设置isDelete= 1, 软删除）
     * @param userId 用户id
     * @param orderId 订单id
     * @return boolean true: 成功; false: 失败
     */
    public boolean deleteOrder(String userId, String orderId);



    public OrderStatusCountsVO getOrderStatusCounts(String userId);


}

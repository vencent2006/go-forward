package com.imooc.service.center;

import com.imooc.pojo.OrderItems;
import com.imooc.pojo.bo.center.OrderItemsCommentBO;
import com.imooc.utils.PagedGridResult;

import java.util.List;

public interface MyCommentsService {

    /**
     * 根据订单id查询关联的商品
     * @param orderId 订单id
     * @return List<OrderItems>
     */
    public List<OrderItems> queryPendingComment(String orderId);


    /**
     * 保存评论
     * @param orderId 订单id
     * @param userId 用户id
     * @param commentList 评论列表
     */
    public void saveComments(String orderId, String userId, List<OrderItemsCommentBO> commentList);


    /**
     * 查询我的评价（分页）
     * @param userId 用户id
     * @param page 第几页
     * @param pageSize 每页大小
     * @return PagedGridResult
     */
    public PagedGridResult queryMyComments(String userId, Integer page, Integer pageSize);

}

package com.imooc.service;

import com.imooc.pojo.Items;
import com.imooc.pojo.ItemsImg;
import com.imooc.pojo.ItemsParam;
import com.imooc.pojo.ItemsSpec;
import com.imooc.pojo.vo.CommentLevelCountsVO;
import com.imooc.pojo.vo.ShopcartVO;
import com.imooc.utils.PagedGridResult;

import java.util.List;

public interface ItemService {
    /**
     * 根据商品ID查询详情
     * @param itemId 商品id（分库分表）
     * @return Items
     */
    public Items queryItemById(String itemId);

    /**
     * 根据商品id查询商品图片列表
     * @param itemId 商品id
     * @return List<ItemsImg>
     */
    public List<ItemsImg> queryItemImgList(String itemId);

    /**
     * 根据商品id查询商品规格
     * @param itemId 商品id
     * @return
     */
    public List<ItemsSpec> queryItemSpecList(String itemId);

    /**
     * 根据商品id查询商品参数
     * @param itemId 商品id
     * @return ItemsParam
     */
    public ItemsParam queryItemParam(String itemId);

    /**
     * 根据商品id查询商品的评价等级数量
     * @param itemId 商品id
     * @return CommentLevelCountsVO
     */
    public CommentLevelCountsVO queryCommentCounts(String itemId);

    /**
     * 根据商品id查询商品的评价（分页）
     * @param itemId 商品id
     * @param level 商品id
     * @param page 第几页
     * @param pageSize 每页多少条
     * @return PagedGridResult
     */
    public PagedGridResult queryPagedComments(String itemId, Integer level, Integer page, Integer pageSize);


    /**
     * 搜索商品列表（分页）
     * @param keywords 关键字
     * @param sort 排序
     * @param page 第几页
     * @param pageSize 每页多少条
     * @return PagedGridResult
     */
    public PagedGridResult searchItems(String keywords, String sort, Integer page, Integer pageSize);


    /**
     * 根据分类搜索商品列表（分页）
     * @param catId 分类id
     * @param sort 排序
     * @param page 第几页
     * @param pageSize 每页多少条
     * @return PagedGridResult
     */
    public PagedGridResult searchItems(Integer catId, String sort, Integer page, Integer pageSize);


    /**
     * 根据规格ids查询最新的购物车中商品数据（用于刷新渲染购物车中的商品数据）
     * @param specIds 规格id（以逗号分割的id）
     * @return List<ShopcartVO>
     */
    public List<ShopcartVO> queryItemsBySpecIds(String specIds);

    /**
     * 根据商品规格id获取规格对象的具体信息
     * @param specId 商品规格id
     * @return ItemsSpec 规格对象
     */
    public ItemsSpec queryItemSpecById(String specId);

    /**
     * 根据商品id获取商品的主图
     * @param itemId 商品id
     * @return String 主图的地址
     */
    public String queryItemMainImgById(String itemId);

    /**
     * 减少库存
     * @param specId 规格id
     * @param buyCounts 减少量
     */
    public void decreaseItemSpecStock(String specId, int buyCounts);
}

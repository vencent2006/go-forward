package com.imooc.service;

import com.imooc.pojo.Items;
import com.imooc.pojo.ItemsImg;
import com.imooc.pojo.ItemsParam;
import com.imooc.pojo.ItemsSpec;

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
}

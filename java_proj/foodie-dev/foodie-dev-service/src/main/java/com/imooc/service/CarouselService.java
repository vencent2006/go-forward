package com.imooc.service;

import com.imooc.pojo.Carousel;

import java.util.List;

public interface CarouselService {
    /**
     * 查询所有轮播图列表
     * @param isShow 是否显示
     * @return List<Carousel>
     */
    public List<Carousel> queryAll(Integer isShow);
}

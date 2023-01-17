package com.imooc.service.impl;

import com.github.pagehelper.PageInfo;
import com.imooc.utils.PagedGridResult;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class BaseService {

    /**
     * 设置分页信息
     * @param list 分页列表
     * @param page 第几页
     * @return PagedGridResult
     */
    protected PagedGridResult setterPagedGrid(List<?> list, Integer page) {
        PageInfo<?> pageList = new PageInfo<>(list);
        PagedGridResult grid = new PagedGridResult();
        grid.setPage(page);// 当前第几页
        grid.setRows(list);// 分页后的数据
        grid.setTotal(pageList.getPages());// 总页数
        grid.setRecords(pageList.getTotal());// 总记录数

        return grid;
    }
}

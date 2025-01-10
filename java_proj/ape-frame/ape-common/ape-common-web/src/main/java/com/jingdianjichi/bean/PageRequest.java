package com.jingdianjichi.bean;

import lombok.Setter;

@Setter
public class PageRequest {
    private Long pageNo = 1L;
    private Long pageSize = 10L;

    public Long getPageNo() {
        if (pageNo == null || pageNo < 1) {
            pageNo = 1L;
        }
        return pageNo;
    }

    public Long getPageSize() {
        if (pageSize == null || pageSize < 1 || pageSize > Integer.MAX_VALUE) {
            pageSize = 10L;
        }
        return pageSize;
    }
}

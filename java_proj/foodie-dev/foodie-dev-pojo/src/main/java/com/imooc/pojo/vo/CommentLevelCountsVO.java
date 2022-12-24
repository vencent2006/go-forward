package com.imooc.pojo.vo;

/**
 * 用于展示商品评价数量的VO
 */
public class CommentLevelCountsVO {
    public Integer totalCounts;     // 总数
    public Integer goodCounts;      // 好评数
    public Integer normalCounts;    // 中评数
    public Integer badCounts;       // 差评数

    public Integer getTotalCounts() {
        return totalCounts;
    }

    public void setTotalCounts(Integer totalCounts) {
        this.totalCounts = totalCounts;
    }

    public Integer getGoodCounts() {
        return goodCounts;
    }

    public void setGoodCounts(Integer goodCounts) {
        this.goodCounts = goodCounts;
    }

    public Integer getNormalCounts() {
        return normalCounts;
    }

    public void setNormalCounts(Integer normalCounts) {
        this.normalCounts = normalCounts;
    }

    public Integer getBadCounts() {
        return badCounts;
    }

    public void setBadCounts(Integer badCounts) {
        this.badCounts = badCounts;
    }
}

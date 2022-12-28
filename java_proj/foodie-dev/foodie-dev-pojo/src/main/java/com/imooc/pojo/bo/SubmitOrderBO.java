package com.imooc.pojo.bo;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

/**
 * 用于创建订单的BO对象
 */
@ApiModel(value = "提交订单对象BO", description = "提交订单对象BO")
public class SubmitOrderBO {

    @ApiModelProperty(value = "用户id", name = "userId", example = "imooc1234", required = true)
    private String userId;
    @ApiModelProperty(value = "规格id", name = "itemSpecIds", example = "bingan-1001-spec-1,bingan-1001-spec-2", required = true)
    private String itemSpecIds;
    @ApiModelProperty(value = "地址id", name = "addressId", example = "2212270KC97T79AW", required = true)
    private String addressId;
    @ApiModelProperty(value = "支付方式", name = "payMethod", example = "1", required = true)
    private Integer payMethod;
    @ApiModelProperty(value = "买家留言", name = "leftMsg", example = "需要包装盒", required = true)
    private String leftMsg;

    public String getUserId() {
        return userId;
    }

    public void setUserId(String userId) {
        this.userId = userId;
    }

    public String getItemSpecIds() {
        return itemSpecIds;
    }

    public void setItemSpecIds(String itemSpecIds) {
        this.itemSpecIds = itemSpecIds;
    }

    public String getAddressId() {
        return addressId;
    }

    public void setAddressId(String addressId) {
        this.addressId = addressId;
    }

    public Integer getPayMethod() {
        return payMethod;
    }

    public void setPayMethod(Integer payMethod) {
        this.payMethod = payMethod;
    }

    public String getLeftMsg() {
        return leftMsg;
    }

    public void setLeftMsg(String leftMsg) {
        this.leftMsg = leftMsg;
    }

    @Override
    public String toString() {
        return "SubmitOrderBO{" +
                "userId='" + userId + '\'' +
                ", itemSpecIds='" + itemSpecIds + '\'' +
                ", addressId='" + addressId + '\'' +
                ", payMethod=" + payMethod +
                ", leftMsg='" + leftMsg + '\'' +
                '}';
    }
}

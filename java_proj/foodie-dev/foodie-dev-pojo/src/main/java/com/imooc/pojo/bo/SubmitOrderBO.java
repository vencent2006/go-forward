package com.imooc.pojo.bo;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

/**
 * 用户新增或修改地址的BO
 */
@ApiModel(value = "收货地址对象BO", description = "收货地址对象BO")
public class AddressBO {
    @ApiModelProperty(value = "地址id", name = "addressId", example = "1", required = false)
    private String addressId;
    @ApiModelProperty(value = "用户id", name = "userId", example = "imooc1234", required = true)
    private String userId;
    @ApiModelProperty(value = "收件人", name = "receiver", example = "张三", required = false)
    private String receiver;
    @ApiModelProperty(value = "电话号码", name = "mobile", example = "13888888888", required = false)
    private String mobile;
    @ApiModelProperty(value = "省", name = "province", example = "辽宁省", required = false)
    private String province;
    @ApiModelProperty(value = "市", name = "city", example = "沈阳市", required = false)
    private String city;
    @ApiModelProperty(value = "区", name = "district", example = "和平区", required = false)
    private String district;
    @ApiModelProperty(value = "详细地址", name = "detail", example = "xx楼xx号", required = false)
    private String detail;

    public String getAddressId() {
        return addressId;
    }

    public void setAddressId(String addressId) {
        this.addressId = addressId;
    }

    public String getUserId() {
        return userId;
    }

    public void setUserId(String userId) {
        this.userId = userId;
    }

    public String getReceiver() {
        return receiver;
    }

    public void setReceiver(String receiver) {
        this.receiver = receiver;
    }

    public String getMobile() {
        return mobile;
    }

    public void setMobile(String mobile) {
        this.mobile = mobile;
    }

    public String getProvince() {
        return province;
    }

    public void setProvince(String province) {
        this.province = province;
    }

    public String getCity() {
        return city;
    }

    public void setCity(String city) {
        this.city = city;
    }

    public String getDistrict() {
        return district;
    }

    public void setDistrict(String district) {
        this.district = district;
    }

    public String getDetail() {
        return detail;
    }

    public void setDetail(String detail) {
        this.detail = detail;
    }
}

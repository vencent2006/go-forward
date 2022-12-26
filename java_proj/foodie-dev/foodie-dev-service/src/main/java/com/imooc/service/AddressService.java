package com.imooc.service;

import com.imooc.pojo.UserAddress;

import java.util.List;

public interface AddressService {

    /**
     * 根据用户id查询用户的收货地址列表
     * @param userId 用户id
     * @return List<UserAddress>
     */
    public List<UserAddress> queryAll(String userId);
}

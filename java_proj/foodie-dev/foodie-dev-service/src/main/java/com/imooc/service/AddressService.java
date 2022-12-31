package com.imooc.service;

import com.imooc.pojo.UserAddress;
import com.imooc.pojo.bo.AddressBO;

import java.util.List;

public interface AddressService {

    /**
     * 根据用户id查询用户的收货地址列表
     * @param userId 用户id
     * @return List<UserAddress>
     */
    public List<UserAddress> queryAll(String userId);

    /**
     * 用户新增地址
     * @param addressBO 新增地址BO
     */
    public void addNewUserAddress(AddressBO addressBO);

    /**
     * 用户修改地址
     * @param addressBO addressId要填写
     */
    public void updateUserAddress(AddressBO addressBO);

    /**
     * 根据用户id和地址id，用户删除地址
     * @param userId 用户id
     * @param addressId 地址id
     */
    public void deleteUserAddress(String userId, String addressId);

    /**
     * 设置用户默认地址
     * @param userId 用户id
     * @param addressId 收货地址id
     */
    public void updateUserAddressToBeDefault(String userId, String addressId);

    /**
     * 根据用户id和地址id，查询具体的用户地址对象信息
     * @param userId 用户id
     * @param addressId 收货地址id
     * @return UserAddress
     */
    public UserAddress queryUserAddress(String userId, String addressId);
}

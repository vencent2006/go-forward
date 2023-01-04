package com.imooc.service.center;

import com.imooc.pojo.Users;
import com.imooc.pojo.bo.center.CenterUserBO;

public interface CenterUserService {
    /**
     * 根据用户id查询用户信息
     * @param userId 用户id
     * @return Users
     */
    public Users queryUserInfo(String userId);


    /**
     * 修改用户信息
     * @param userId 用户id
     * @param centerUserBO 客户端提交的修改信息
     * @return Users
     */
    public Users updateUserInfo(String userId, CenterUserBO centerUserBO);
}

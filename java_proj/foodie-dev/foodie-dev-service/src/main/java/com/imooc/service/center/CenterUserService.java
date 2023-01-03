package com.imooc.service.center;

import com.imooc.pojo.Users;

public interface CenterUserService {
    /**
     * 根据用户id查询用户信息
     * @param userId 用户id
     * @return Users
     */
    public Users queryUserInfo(String userId);
}

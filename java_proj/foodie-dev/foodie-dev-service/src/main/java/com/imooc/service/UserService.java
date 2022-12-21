package com.imooc.service;

import com.imooc.pojo.Users;
import com.imooc.pojo.bo.UserBO;

public interface UserService {
    /**
     * 用户名是否存在
     * @param username 目标用户名
     * @return boolean true:已经存在;false,不存在
     */
    public boolean queryUsernameIsExist(String username);

    /**
     * 创建用户
     * @param userBO 创建用户的客户端参数
     * @return Users
     */
    public Users createUser(UserBO userBO);

    /**
     * 检索用户名和密码是否匹配，用于登录
     * @param username 用户名
     * @param password 密码
     * @return Users
     */
    public Users queryUserForLogin(String username, String password);

}

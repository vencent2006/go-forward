package com.imooc.service;

public interface UserService {
    /**
     * 用户名是否存在
     * @param username 目标用户名
     * @return boolean true:已经存在;false,不存在
     */
    public boolean queryUsernameIsExist(String username);
}

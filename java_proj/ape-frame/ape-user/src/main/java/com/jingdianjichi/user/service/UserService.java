package com.jingdianjichi.user.service;

import com.jingdianjichi.user.entity.dto.UserDto;

// controller -> service 用dto
// service -> mapper 用entity(po)
public interface UserService {
    int addUser(UserDto userDto);

    int delete(Integer id);
}

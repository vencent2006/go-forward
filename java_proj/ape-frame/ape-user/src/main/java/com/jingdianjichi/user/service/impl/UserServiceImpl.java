package com.jingdianjichi.user.service.impl;

import com.jingdianjichi.user.entity.dto.UserDto;
import com.jingdianjichi.user.entity.po.UserPo;
import com.jingdianjichi.user.mapper.UserMapper;
import com.jingdianjichi.user.service.UserService;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserServiceImpl implements UserService {
    @Autowired
    private UserMapper userMapper;

    @Override
    public int addUser(UserDto userDto) {
        UserPo userPo = new UserPo();
        BeanUtils.copyProperties(userDto, userPo);
        int insert = userMapper.insert(userPo);
        return insert;
    }

    @Override
    public int delete(Integer id) {
        return userMapper.deleteById(id);
    }
}

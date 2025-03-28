package com.jingdianjichi.user.dao;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.baomidou.mybatisplus.core.metadata.IPage;
import com.jingdianjichi.user.entity.po.UserPo;
import org.springframework.stereotype.Repository;

@Repository
public interface UserDao extends BaseMapper<UserPo> {
    IPage<UserPo> getUserPage(IPage<UserPo> userPoPage);
}

package com.jingdianjichi.user.convert;

import com.jingdianjichi.user.entity.po.SysUser;
import com.jingdianjichi.user.entity.req.SysUserReq;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

@Mapper
public interface SysUserConverter {
    // 单例
    SysUserConverter INSTANCE = Mappers.getMapper(SysUserConverter.class);

    // Mapping注解，可以修改字段名
    // @Mapping(source = "pageNo", target = "pageStart")
    SysUser convertReqToSysUser(SysUserReq sysUserReq);
}

package com.vs.myworld.business.mapper;

import com.vs.myworld.business.domain.SmsCode;
import com.vs.myworld.business.domain.SmsCodeExample;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface SmsCodeMapper {
    long countByExample(SmsCodeExample example);

    int deleteByExample(SmsCodeExample example);

    int deleteByPrimaryKey(Long id);

    int insert(SmsCode row);

    int insertSelective(SmsCode row);

    List<SmsCode> selectByExample(SmsCodeExample example);

    SmsCode selectByPrimaryKey(Long id);

    int updateByExampleSelective(@Param("row") SmsCode row, @Param("example") SmsCodeExample example);

    int updateByExample(@Param("row") SmsCode row, @Param("example") SmsCodeExample example);

    int updateByPrimaryKeySelective(SmsCode row);

    int updateByPrimaryKey(SmsCode row);
}
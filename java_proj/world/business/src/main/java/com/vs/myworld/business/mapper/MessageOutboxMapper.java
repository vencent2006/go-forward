package com.vs.myworld.business.mapper;

import com.vs.myworld.business.domain.MessageOutbox;
import com.vs.myworld.business.domain.MessageOutboxExample;
import java.util.List;
import org.apache.ibatis.annotations.Param;

public interface MessageOutboxMapper {
    long countByExample(MessageOutboxExample example);

    int deleteByExample(MessageOutboxExample example);

    int deleteByPrimaryKey(Long id);

    int insert(MessageOutbox row);

    int insertSelective(MessageOutbox row);

    List<MessageOutbox> selectByExample(MessageOutboxExample example);

    MessageOutbox selectByPrimaryKey(Long id);

    int updateByExampleSelective(@Param("row") MessageOutbox row, @Param("example") MessageOutboxExample example);

    int updateByExample(@Param("row") MessageOutbox row, @Param("example") MessageOutboxExample example);

    int updateByPrimaryKeySelective(MessageOutbox row);

    int updateByPrimaryKey(MessageOutbox row);
}
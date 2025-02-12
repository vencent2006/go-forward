package com.vs.myworld.business.mapper;

import com.vs.myworld.business.domain.MessageInbox;
import com.vs.myworld.business.domain.MessageInboxExample;
import java.util.List;
import org.apache.ibatis.annotations.Param;

public interface MessageInboxMapper {
    long countByExample(MessageInboxExample example);

    int deleteByExample(MessageInboxExample example);

    int deleteByPrimaryKey(Long id);

    int insert(MessageInbox row);

    int insertSelective(MessageInbox row);

    List<MessageInbox> selectByExample(MessageInboxExample example);

    MessageInbox selectByPrimaryKey(Long id);

    int updateByExampleSelective(@Param("row") MessageInbox row, @Param("example") MessageInboxExample example);

    int updateByExample(@Param("row") MessageInbox row, @Param("example") MessageInboxExample example);

    int updateByPrimaryKeySelective(MessageInbox row);

    int updateByPrimaryKey(MessageInbox row);
}
package com.vs.myworld.business.mapper;

import com.vs.myworld.business.domain.WorldEvent;
import com.vs.myworld.business.domain.WorldEventExample;
import java.util.List;
import org.apache.ibatis.annotations.Param;

public interface WorldEventMapper {
    long countByExample(WorldEventExample example);

    int deleteByExample(WorldEventExample example);

    int deleteByPrimaryKey(Long id);

    int insert(WorldEvent row);

    int insertSelective(WorldEvent row);

    List<WorldEvent> selectByExample(WorldEventExample example);

    WorldEvent selectByPrimaryKey(Long id);

    int updateByExampleSelective(@Param("row") WorldEvent row, @Param("example") WorldEventExample example);

    int updateByExample(@Param("row") WorldEvent row, @Param("example") WorldEventExample example);

    int updateByPrimaryKeySelective(WorldEvent row);

    int updateByPrimaryKey(WorldEvent row);
}
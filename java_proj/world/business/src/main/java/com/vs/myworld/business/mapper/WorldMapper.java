package com.vs.myworld.business.mapper;

import com.vs.myworld.business.domain.World;
import com.vs.myworld.business.domain.WorldExample;
import java.util.List;
import org.apache.ibatis.annotations.Param;

public interface WorldMapper {
    long countByExample(WorldExample example);

    int deleteByExample(WorldExample example);

    int deleteByPrimaryKey(Long id);

    int insert(World row);

    int insertSelective(World row);

    List<World> selectByExample(WorldExample example);

    World selectByPrimaryKey(Long id);

    int updateByExampleSelective(@Param("row") World row, @Param("example") WorldExample example);

    int updateByExample(@Param("row") World row, @Param("example") WorldExample example);

    int updateByPrimaryKeySelective(World row);

    int updateByPrimaryKey(World row);
}
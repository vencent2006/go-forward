package com.vs.myworld.business.mapper;

import com.vs.myworld.business.domain.WorldUserRelation;
import com.vs.myworld.business.domain.WorldUserRelationExample;
import java.util.List;
import org.apache.ibatis.annotations.Param;

public interface WorldUserRelationMapper {
    long countByExample(WorldUserRelationExample example);

    int deleteByExample(WorldUserRelationExample example);

    int deleteByPrimaryKey(Long id);

    int insert(WorldUserRelation row);

    int insertSelective(WorldUserRelation row);

    List<WorldUserRelation> selectByExample(WorldUserRelationExample example);

    WorldUserRelation selectByPrimaryKey(Long id);

    int updateByExampleSelective(@Param("row") WorldUserRelation row, @Param("example") WorldUserRelationExample example);

    int updateByExample(@Param("row") WorldUserRelation row, @Param("example") WorldUserRelationExample example);

    int updateByPrimaryKeySelective(WorldUserRelation row);

    int updateByPrimaryKey(WorldUserRelation row);
}
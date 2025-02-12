package com.vs.myworld.business.mapper;

import com.vs.myworld.business.domain.WorldCharacterRelation;
import com.vs.myworld.business.domain.WorldCharacterRelationExample;
import java.util.List;
import org.apache.ibatis.annotations.Param;

public interface WorldCharacterRelationMapper {
    long countByExample(WorldCharacterRelationExample example);

    int deleteByExample(WorldCharacterRelationExample example);

    int deleteByPrimaryKey(Long id);

    int insert(WorldCharacterRelation row);

    int insertSelective(WorldCharacterRelation row);

    List<WorldCharacterRelation> selectByExample(WorldCharacterRelationExample example);

    WorldCharacterRelation selectByPrimaryKey(Long id);

    int updateByExampleSelective(@Param("row") WorldCharacterRelation row, @Param("example") WorldCharacterRelationExample example);

    int updateByExample(@Param("row") WorldCharacterRelation row, @Param("example") WorldCharacterRelationExample example);

    int updateByPrimaryKeySelective(WorldCharacterRelation row);

    int updateByPrimaryKey(WorldCharacterRelation row);
}
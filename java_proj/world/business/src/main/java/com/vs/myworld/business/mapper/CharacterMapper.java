package com.vs.myworld.business.mapper;

import com.vs.myworld.business.domain.Character;
import com.vs.myworld.business.domain.CharacterExample;
import java.util.List;
import org.apache.ibatis.annotations.Param;

public interface CharacterMapper {
    long countByExample(CharacterExample example);

    int deleteByExample(CharacterExample example);

    int deleteByPrimaryKey(Long id);

    int insert(Character row);

    int insertSelective(Character row);

    List<Character> selectByExample(CharacterExample example);

    Character selectByPrimaryKey(Long id);

    int updateByExampleSelective(@Param("row") Character row, @Param("example") CharacterExample example);

    int updateByExample(@Param("row") Character row, @Param("example") CharacterExample example);

    int updateByPrimaryKeySelective(Character row);

    int updateByPrimaryKey(Character row);
}
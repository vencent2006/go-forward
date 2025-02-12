package com.vs.myworld.business.mapper;

import com.vs.myworld.business.domain.File;
import com.vs.myworld.business.domain.FileExample;
import java.util.List;
import org.apache.ibatis.annotations.Param;

public interface FileMapper {
    long countByExample(FileExample example);

    int deleteByExample(FileExample example);

    int deleteByPrimaryKey(Long id);

    int insert(File row);

    int insertSelective(File row);

    List<File> selectByExample(FileExample example);

    File selectByPrimaryKey(Long id);

    int updateByExampleSelective(@Param("row") File row, @Param("example") FileExample example);

    int updateByExample(@Param("row") File row, @Param("example") FileExample example);

    int updateByPrimaryKeySelective(File row);

    int updateByPrimaryKey(File row);
}
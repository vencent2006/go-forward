package com.vincent.mapper;

import com.vincent.pojo.DbStu;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface DbStuMapperCustom {
    public List<DbStu> getStuById(String sid);
}

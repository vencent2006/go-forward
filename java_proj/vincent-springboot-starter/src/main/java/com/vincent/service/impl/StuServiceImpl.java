package com.vincent.service.impl;

import com.vincent.mapper.DbStuMapper;
import com.vincent.pojo.DbStu;
import com.vincent.service.StuService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import tk.mybatis.mapper.entity.Example;

import java.util.List;

@Service
public class StuServiceImpl implements StuService {
    @Autowired
    private DbStuMapper stuMapper;

    @Override
    public void saveStu(DbStu stu) {
        stuMapper.insert(stu);
    }

    @Override
    public DbStu queryById(String id) {
        DbStu stu = stuMapper.selectByPrimaryKey(id);
        return stu;
    }

    @Override
    public List<DbStu> queryByCondition(String name, Integer sex) {
        Example example = new Example(DbStu.class);
        Example.Criteria criteria = example.createCriteria();
        criteria.andEqualTo("name", name);
        criteria.andEqualTo("sex", sex);
        List<DbStu> list = stuMapper.selectByExample(example);
        return list;
    }


    // @Override
    // public List<DbStu> queryByCondition(String name, Integer sex) {
    //     DbStu stu = new DbStu();
    //     stu.setName(name);
    //     stu.setSex(sex);
    //     List<DbStu> list = stuMapper.select(stu);
    //     return list;
    // }
}

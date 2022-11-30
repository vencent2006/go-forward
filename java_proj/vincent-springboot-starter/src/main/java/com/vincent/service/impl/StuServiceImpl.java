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


    @Override
    public void updateStu(DbStu stu) {
        stuMapper.updateByPrimaryKey(stu);
    }

    @Override
    public void deleteStu(DbStu stu) {
        // 删除对象/数据的三种方式
        // 1. 根据主键删除
        // stuMapper.deleteByPrimaryKey(stu);
        // 2. 根据对象中的属性值匹配做条件删除
        // stuMapper.delete(stu);
        // 3. 根据构建的example进行条件删除
        Example example = new Example(DbStu.class);
        Example.Criteria criteria = example.createCriteria();
        criteria.andEqualTo("name", stu.getName());
        stuMapper.deleteByExample(example);
    }

}

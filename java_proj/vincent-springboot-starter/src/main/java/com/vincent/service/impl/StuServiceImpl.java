package com.vincent.service.impl;

import com.vincent.mapper.DbStuMapper;
import com.vincent.mapper.DbStuMapperCustom;
import com.vincent.pojo.DbStu;
import com.vincent.service.StuService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;
import tk.mybatis.mapper.entity.Example;

import java.util.List;
import java.util.UUID;

@Service
public class StuServiceImpl implements StuService {
    @Autowired
    private DbStuMapper stuMapper;

    @Autowired
    private DbStuMapperCustom stuMapperCustom;

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
    public DbStu queryByIdCustom(String id) {
        // try {
        //     Thread.sleep(3500);
        // } catch (InterruptedException e) {
        //     e.printStackTrace();
        // }

        List<DbStu> list = stuMapperCustom.getStuById(id);
        if (list != null && !list.isEmpty()){
            return list.get(0);
        }
        return null;
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

    @Transactional(propagation = Propagation.REQUIRED)
    @Override
    public void testTransaction() {
        // 1. 新增数据
        String sid = UUID.randomUUID().toString();
        DbStu stu = new DbStu();
        stu.setId(sid);
        stu.setName(sid);
        stu.setSex(2);

        stuMapper.insert(stu);

        // 异常
        int a = 100/0;

        // 2. 修改数据
        DbStu stuUpdate = new DbStu();
        stuUpdate.setId("1001");
        stuUpdate.setName("1001修改");
        stuUpdate.setSex(3);
        stuMapper.updateByPrimaryKeySelective(stuUpdate);// 只对修改的值进行修改


        // 3. 模拟发生异常
        // 4. 观察1和2步骤所发生的数据变动，有没有影响到数据库
        // 5. 处理事务，实现事务回滚，不让先前的数据入库


    }
}

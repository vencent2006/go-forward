package com.vincent.service.impl;

import com.vincent.mapper.DbStuMapper;
import com.vincent.pojo.DbStu;
import com.vincent.service.StuService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class StuServiceImpl implements StuService {
    @Autowired
    private DbStuMapper stuMapper;

    @Override
    public void saveStu(DbStu stu) {
        stuMapper.insert(stu);
    }
}

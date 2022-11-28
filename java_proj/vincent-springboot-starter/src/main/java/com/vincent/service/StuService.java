package com.vincent.service;

import com.vincent.pojo.DbStu;

public interface StuService {
    /**
     * 新增stu到数据库
     * @param stu
     */
    public void saveStu(DbStu stu);
}

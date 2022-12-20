package com.test;

import com.imooc.Application;
import com.imooc.pojo.Stu;
import com.imooc.service.StuService;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

@RunWith(SpringRunner.class)
@SpringBootTest(classes = Application.class)
public class TransTest {
    @Autowired
    private StuService stuService;

    @Test
    public void myTest() {
        Stu stuInfo = stuService.getStuInfo(1001);
        System.out.println(stuInfo.getId());
        System.out.println(stuInfo.getName());
        System.out.println(stuInfo.getAge());
    }
}

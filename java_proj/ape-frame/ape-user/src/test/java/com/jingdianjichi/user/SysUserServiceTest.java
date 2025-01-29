package com.jingdianjichi.user;

import com.jingdianjichi.user.entity.po.SysUser;
import com.jingdianjichi.user.service.SysUserService;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import javax.annotation.Resource;

// classes 表示启动类
// webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT 表示随机端口启动应用
@SpringBootTest(classes = UserApplication.class, webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@RunWith(SpringRunner.class) // 表示启动这个单元测试类（单元测试类是不能够运行的），需要传递一个参数，必须是SpringRunner的实例类型
public class SysUserServiceTest {
    @Resource
    private SysUserService sysUserService;

    @Test
    public void testQuery() {
        SysUser sysUser = sysUserService.queryById(1L);
        System.out.println(sysUser);
    }
}

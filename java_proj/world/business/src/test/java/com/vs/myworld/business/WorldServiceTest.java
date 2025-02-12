package com.vs.myworld.business;

import com.vs.myworld.business.config.BusinessApplication;
import com.vs.myworld.business.resp.WorldQueryResp;
import com.vs.myworld.business.service.WorldService;
import com.vs.myworld.business.util.PrintUtil;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import javax.annotation.Resource;
import java.util.List;

@SpringBootTest(classes = BusinessApplication.class, webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@RunWith(SpringRunner.class) // 表示启动这个单元测试类（单元测试类是不能够运行的），需要传递一个参数，必须是SpringRunner的实例类型
public class WorldServiceTest {
    @Resource
    private WorldService worldService;

    @Test
    public void testQuery() {
        List<WorldQueryResp> list = worldService.query();
        // json的格式化打印
        PrintUtil.printJsonFormat(list);
    }
}

package com.vincent.test;

import com.vincent.huobi.client.MarketClient;
import com.vincent.huobi.client.req.market.MarketDetailMergedRequest;
import com.vincent.huobi.constant.HuobiOptions;
import com.vincent.huobi.model.market.MarketDetailMerged;
import com.vincent.pojo.DbStu;
import com.vincent.service.StuService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.UUID;

@SpringBootTest // 表示当前会被springboot加载，加载为测试类
public class MyTest {
    @Autowired
    private StuService stuService;

    @Test
    public void testSaveStu() {
        DbStu stu = new DbStu();
        stu.setId(UUID.randomUUID().toString());
        stu.setName("test");
        stu.setSex(1);
        stuService.saveStu(stu);
        // int a = 1/0;
    }

    @Test
    public void testHuobi() {
        MarketClient marketClient = MarketClient.create(new HuobiOptions());
        String symbol = "btcusdt";

        MarketDetailMerged marketDetailMerged = marketClient.getMarketDetailMerged(MarketDetailMergedRequest.builder().symbol(symbol).build());
        System.out.println(marketDetailMerged.toString());
    }

}


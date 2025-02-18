package com.vs.myworld.business.cardlist;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import org.junit.Test;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

@SpringBootTest
public class CardListTest {
    @Test
    public void test() throws JsonProcessingException {
        int fansNum = 5000; // 假设粉丝数（根据需要修改）

        // 1. 创建 EnterInfo 对象
        EnterInfo weiboEnter = new EnterInfo();
        weiboEnter.setTitle("微博打赏");
        weiboEnter.setIcon("");
        weiboEnter.setDesc("大于1万粉丝可开启");
        weiboEnter.setIsSatisfy(fansNum >= 10000 ? 1 : 0);
        weiboEnter.setCertificationTitle(fansNum >= 10000 ? "已满足" : "未满足");
        weiboEnter.setCertificationLink("");

        EnterInfo articleEnter = new EnterInfo();
        articleEnter.setTitle("文章打赏");
        articleEnter.setIcon("");
        articleEnter.setDesc("大于1千粉丝可开启");
        articleEnter.setIsSatisfy(fansNum >= 1000 ? 1 : 0);
        articleEnter.setCertificationTitle(fansNum >= 1000 ? "已满足" : "未满足");
        articleEnter.setCertificationLink("");

        // 2. 创建 Settle 对象
        Settle settle = new Settle();
        settle.setSettleTitle("付费打赏");
        settle.setSettleType("tip");
        settle.setSettleColour(1);
        settle.setSettleDesc("在微博或文章中获得粉丝付费打赏");
        settle.setSettleBtnText(fansNum >= 1000 ? (fansNum >= 10000 ? "立即开启" : "开启文章打赏") : "立即开启");
        settle.setSettleLink("//paid.e.weibo.com/v1/public/h5/contentincome/rewardsetting");
        settle.setSettleBtnTextStatue(fansNum >= 1000 ? 1 : 0);
        settle.setRecruitInfo(new ArrayList<>());
        settle.setEnterInfo(Arrays.asList(weiboEnter, articleEnter));

        // 3. 创建 CardStatus 对象
        CardStatus cardStatus = new CardStatus();
        cardStatus.setCardStatus(3);
        cardStatus.setTitle("创收工具");
        cardStatus.setDesc("");
        cardStatus.setSatisfyShow(0);
        cardStatus.setSatisfyStatue(1);
        cardStatus.setRecommendText("");
        cardStatus.setSettle(Collections.singletonList(settle));

        ObjectMapper mapper = new ObjectMapper();
        mapper.enable(SerializationFeature.INDENT_OUTPUT);
        String json = mapper.writeValueAsString(cardStatus);
        System.out.println(json);
    }
}

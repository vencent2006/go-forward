package com.jingdianjichi.user;

import com.jingdianjichi.redis.util.RedisUtil;
import lombok.extern.slf4j.Slf4j;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import javax.annotation.Resource;

/**
 * Redis的lua测试类
 *
 * @author: ChickenWing
 * @date: 2023/1/16
 */
@SpringBootTest(classes = UserApplication.class, webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@RunWith(SpringRunner.class)
@Slf4j
public class RedisUtilTest {

    @Resource
    private RedisUtil redisUtil;

    @Test
    public void testCompareAndSet() throws Exception {
        Boolean result = redisUtil.compareAndSet("luaCas", 2L, 3L);
        log.info("RedisUtilTest.testCompareAndSet.result:{}", result);
    }

}

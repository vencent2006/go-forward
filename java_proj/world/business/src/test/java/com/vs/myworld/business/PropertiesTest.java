package com.vs.myworld.business;

import com.vs.myworld.business.common.OtherProperties;
import com.vs.myworld.business.common.OtherYml;
import com.vs.myworld.business.config.BusinessApplication;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import javax.annotation.Resource;

@RunWith(SpringRunner.class)
@SpringBootTest(classes = BusinessApplication.class, webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class PropertiesTest {
    @Value("${spring.datasource.url}")
    private String dataUrl;

    @Resource
    private OtherProperties otherProperties;

    @Resource
    private OtherYml otherYml;

    @Test
    public void testReadPort() {
        System.out.println("spring.datasource.url: " + dataUrl);
    }

    @Test
    public void testReadOtherProperties() {
        System.out.println("other.title: " + otherProperties.getTitle());
        System.out.println("other.blog: " + otherProperties.getBlog());
    }

    @Test
    public void testReadOtherYml() {
        System.out.println("other.title: " + otherYml.getTitle());
        System.out.println("other.blog: " + otherYml.getBlog());
    }
}

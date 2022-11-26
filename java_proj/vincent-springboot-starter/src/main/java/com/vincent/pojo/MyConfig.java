package com.vincent.pojo;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.PropertySource;
import org.springframework.stereotype.Component;

@Component // 把本配置放到springboot容器中，使其被扫描到
@ConfigurationProperties(prefix = "user")
@PropertySource(value = "classpath:MyConfig.properties", encoding = "utf-8")
public class MyConfig {
    public String uname;
    public Integer age;
    public String sex;

    // 绑定了properties，不需要构造函数

    public String getUname() {
        return uname;
    }

    public void setUname(String uname) {
        this.uname = uname;
    }

    public Integer getAge() {
        return age;
    }

    public void setAge(Integer age) {
        this.age = age;
    }

    public String getSex() {
        return sex;
    }

    public void setSex(String sex) {
        this.sex = sex;
    }
}

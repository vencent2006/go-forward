package com.vincent.controller;

import com.vincent.pojo.MyConfig;
import com.vincent.pojo.Stu;
import com.vincent.pojo.Student;
import lombok.extern.slf4j.Slf4j;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@Slf4j
public class HelloController {
    private static final Logger logger = LoggerFactory.getLogger("name");


    @GetMapping("hello")
    public String hello(){
        logger.info("hello world------------------------");
        return "hello world";
    }


    // 直接从Bean里读取
    @Autowired
    public Stu stu;
    @GetMapping("stu")
    public Object getStu(){
        return stu;
    }

    // 从properties里读取
    @Autowired
    public MyConfig myConfig;
    @GetMapping("myconfig")
    public Object myConfig(){
        return myConfig;
    }

    // 从application.yml中读取
    @Value("${self.custom.config.sdkSecret}")
    public String sdkSecret;
    @Value("${self.custom.config.host}")
    public String host;
    @Value("${self.custom.config.port}")
    public String port;
    @Value("${app.name.xxx.yyy.zzz}")
    public String xyz;
    @GetMapping("yml_config")
    public Object config(){
        return sdkSecret + ";\t" + host + ";\t" + xyz + ";\t" + port;
    }

    // lombok annotation
    // @Autowired
    public Student student;
    @GetMapping("student")
    public Object getStudent(){
        Student stu = new Student();
        stu.setAge(11);
        stu.setName("michael");

        log.debug(stu.toString());
        log.info(stu.toString());
        log.warn(stu.toString());
        log.error(stu.toString());

        return stu;
    }
}

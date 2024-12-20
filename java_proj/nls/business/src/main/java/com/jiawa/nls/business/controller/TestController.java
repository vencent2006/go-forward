package com.jiawa.nls.business.controller;

import com.jiawa.nls.business.domain.Demo;
import com.jiawa.nls.business.service.DemoService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.util.List;

@RestController
public class TestController {
    @Resource
    private DemoService demoService;

    @GetMapping("/hello1")
    public String hello(){
        return "Hello world!";
    }

    @GetMapping("/count")
    public int count(){
        return demoService.count();
    }
    @GetMapping("/query")
    public List<Demo> query(String mobile){
        return demoService.query(mobile);
    }
}

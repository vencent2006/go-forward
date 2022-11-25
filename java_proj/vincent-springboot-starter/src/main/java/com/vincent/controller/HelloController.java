package com.vincent.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HelloController {
    private static final Logger logger = LoggerFactory.getLogger("name");
    @GetMapping("hello")
    public String hello(){
        logger.info("hello world------------------------");
        return "hello world";
    }
}

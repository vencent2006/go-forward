package com.jiawa.nls.business.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TestController {
    @GetMapping("/hello")
    private String hello(){
        return "Hello World\n";
    }

}

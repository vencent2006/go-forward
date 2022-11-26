package com.vincent.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("stu")
@Slf4j
public class StuController {
    @GetMapping("get")
    public Object get(){
        return "查询";
    }

    @PostMapping("post")
    public Object post(){
        return "修改";
    }

    @PutMapping("put")
    public Object put(){
        return "插入";
    }

    @DeleteMapping("delete")
    public Object delete(){
        return "删除";
    }
}

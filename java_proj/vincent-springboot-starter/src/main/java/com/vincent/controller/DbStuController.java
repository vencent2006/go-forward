package com.vincent.controller;

import com.vincent.pojo.DbStu;
import com.vincent.pojo.bo.DbStuBO;
import com.vincent.service.StuService;
import com.vincent.utils.JSONResult;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.Map;
import java.util.UUID;

@RestController
@RequestMapping("dbstu")
@Slf4j
public class DbStuController {

    @Autowired
    private StuService stuService;


    @GetMapping("{stuId}/get")
    public Object get(@PathVariable("stuId") String stuId,
                      @RequestParam Integer id,
                      @RequestParam String name){
        /**
         * @RequestParam 用于获取url中的请求参数，如果参数变量名保持一致，该注解可以省略
         */
        log.info("stuId = " + stuId);
        log.info("id = " + id);
        log.info("name = " + name);

        return "查询";
    }

    @PostMapping("create")
    public JSONResult createStu(){
        String sid = UUID.randomUUID().toString();
        DbStu stu = new DbStu();
        stu.setId(sid);
        stu.setName("mike");
        stu.setSex(1);

        stuService.saveStu(stu);

        return JSONResult.ok();
    }

    @PostMapping("create2")
    public JSONResult createStu2(@RequestBody DbStuBO stuBO){

        DbStu stu = new DbStu();
        BeanUtils.copyProperties(stuBO, stu);
        String sid = UUID.randomUUID().toString();
        stu.setId(sid);

        stuService.saveStu(stu);

        return JSONResult.ok();
    }
}

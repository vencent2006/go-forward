package com.vincent.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.Map;

@RestController
@RequestMapping("stu")
@Slf4j
public class StuController {
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

    @PostMapping("post")
    public Object post(@RequestBody Map<String, Object> map,
                       @RequestHeader("token") String token,
                       @CookieValue("clientId") String clientId,
                       HttpServletRequest request){
        log.info("map = " + map.toString());
        log.info("token = " + token);
        log.info("clientId = " + clientId);

        String tokenFromRequest = request.getHeader("token");
        log.info("tokenFromRequest = " + tokenFromRequest);

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

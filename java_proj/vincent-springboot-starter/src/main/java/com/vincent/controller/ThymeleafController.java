package com.vincent.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;
import org.thymeleaf.TemplateEngine;
import org.thymeleaf.context.Context;

import javax.servlet.http.HttpServletRequest;
import java.io.FileWriter;
import java.io.IOException;
import java.io.Writer;
import java.util.*;

@Controller
@Slf4j
@RequestMapping("thyme")
public class ThymeleafController {

    @GetMapping("hello")
    public String hello(Model model, HttpServletRequest request){
        String stranger = "jack";
        model.addAttribute("there", stranger);

        Date birthday = new Date();
        model.addAttribute("birthday", birthday);

        Integer sex = 2;// 0 female; 1 male; 2 unknown
        model.addAttribute("sex", sex);

        List<String> myList = new ArrayList<>();
        myList.add("Java");
        myList.add("HTML");
        myList.add("大数据");
        myList.add("SpringBoot");
        myList.add("微服务");

        Map<String, Object> myMap = new HashMap<>();
        myMap.put("id", 10010);
        myMap.put("name", "Jack");
        myMap.put("sex", "男");
        myMap.put("loves", myList);

        model.addAttribute("myList", myList);
        model.addAttribute("myMap", myMap);

        // // HttpServletRequest
        // request.setAttribute("englishName", "xyzabc");
        // request.getSession().setAttribute("userToken", "bala-bala-alla-lala");



        return "teacher";
    }

    @Autowired
    private TemplateEngine templateEngine;


    @GetMapping("createHTML")
    @ResponseBody
    public String createHTML(Model model, HttpServletRequest request) throws Throwable {
        String stranger = "jack";
        model.addAttribute("there", stranger);

        Date birthday = new Date();
        model.addAttribute("birthday", birthday);

        Integer sex = 2;// 0 female; 1 male; 2 unknown
        model.addAttribute("sex", sex);

        List<String> myList = new ArrayList<>();
        myList.add("Java");
        myList.add("HTML");
        myList.add("大数据");
        myList.add("SpringBoot");
        myList.add("微服务");

        Map<String, Object> myMap = new HashMap<>();
        myMap.put("id", 10010);
        myMap.put("name", "Jack");
        myMap.put("sex", "男");
        myMap.put("loves", myList);

        model.addAttribute("myList", myList);
        model.addAttribute("myMap", myMap);

        // HttpServletRequest
        request.setAttribute("englishName", "xyzabc");
        request.getSession().setAttribute("userToken", "bala-bala-alla-lala");


        // 开始静态化
        Context context = new Context();
        context.setVariable("birthday", birthday);
        context.setVariable("there", stranger);
        context.setVariable("sex", sex);
        context.setVariable("myList", myList);
        context.setVariable("myMap", myMap);

        // 定义html最终保存生成的目录位置
        String path = "/temp";
        Writer out = new FileWriter(path + "/teacher.html");
        templateEngine.process("teacher", context, out);
        out.close();


        return "ok";
    }
}

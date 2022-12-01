package com.vincent.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.*;

@Controller
@Slf4j
@RequestMapping("thyme")
public class ThymeleafController {

    @GetMapping("hello")
    public String hello(Model model){
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



        return "teacher";
    }
}

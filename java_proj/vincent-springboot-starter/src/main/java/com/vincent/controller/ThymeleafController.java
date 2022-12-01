package com.vincent.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.Date;
import java.util.Map;

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

        return "teacher";
    }
}

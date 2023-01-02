package com.imooc.controller;

import org.springframework.web.bind.annotation.RestController;

@RestController
public class BaseController {
    // 切换大小写：command + shift + u
    public static final Integer COMMENT_PAGE_SIZE = 10;
    public static final Integer PAGE_SIZE = 20;

    public static final String FOODIE_SHOPCAT = "shopcart"; // 前端购物车的cookie的key

}

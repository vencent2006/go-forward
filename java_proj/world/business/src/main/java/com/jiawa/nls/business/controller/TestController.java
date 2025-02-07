package com.jiawa.nls.business.controller;

import com.jiawa.nls.business.req.DemoQueryReq;
import com.jiawa.nls.business.resp.CommonResp;
import com.jiawa.nls.business.resp.DemoQueryResp;
import com.jiawa.nls.business.service.DemoService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import javax.validation.Valid;
import java.util.List;

@RestController
public class TestController {
    @Resource
    private DemoService demoService;

    @GetMapping("/hello1")
    public CommonResp<String> hello() {
        // return "Hello world!";
        return new CommonResp<>("Hello world!");
    }

    @GetMapping("/count")
    public CommonResp<Integer> count() {
        // return demoService.count();
        return new CommonResp<>(demoService.count());
    }

    @GetMapping("/query")
    public CommonResp<List<DemoQueryResp>> query(@Valid DemoQueryReq req) {
        List<DemoQueryResp> demoList = demoService.query(req);
        // CommonResp<List<Demo>> listCommonResp = new CommonResp<>();
        // listCommonResp.setContent(demoList);
        // return listCommonResp;
        return new CommonResp<>(demoList);
    }
}

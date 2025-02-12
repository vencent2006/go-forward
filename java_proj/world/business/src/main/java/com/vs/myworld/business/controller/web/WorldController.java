package com.vs.myworld.business.controller.web;

import com.vs.myworld.business.resp.CommonResp;
import com.vs.myworld.business.resp.WorldQueryResp;
import com.vs.myworld.business.service.WorldService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.util.List;

@Slf4j
@RestController
@RequestMapping("/api/world")
public class WorldController {
    @Resource
    private WorldService worldService;

    @RequestMapping()
    public CommonResp<List<WorldQueryResp>> query() {
        List<WorldQueryResp> list = worldService.query();
        return new CommonResp<>(list);
    }

}

package com.vs.myworld.business.controller.web;

import com.vs.myworld.business.resp.CommonResp;
import com.vs.myworld.business.resp.WorldQueryInfo;
import com.vs.myworld.business.service.WorldService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;

@Slf4j
@RestController
@RequestMapping("/api/world")
public class WorldController {
    @Resource
    private WorldService worldService;

    @GetMapping()
    public CommonResp<WorldQueryInfo> query(@RequestParam(defaultValue = "0") int page,
                                            @RequestParam(defaultValue = "10") int size) {
        WorldQueryInfo info = new WorldQueryInfo();
        info.setList(worldService.query());
        info.setTotal((long) info.getList().size());
        return new CommonResp<>(info);
    }

}
